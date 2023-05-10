package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/abnervoynich/night-bird/app/models"
	"io"
	"log"
	"net/http"
	"reflect"
	"sort"
	"strings"
)

type ContractValidator struct {
	PelletsFlows []models.PelletFlow
	Pellets      []models.Pellet
}

func NewContractValidatorService(pellets []models.Pellet, flows []models.PelletFlow) *ContractValidator {
	return &ContractValidator{
		Pellets:      pellets,
		PelletsFlows: flows,
	}
}

// StartAPIResponsesValidation run asyncronous api calls and validates
func (c *ContractValidator) StartAPIResponsesValidation() error {
	pelletsMap := c.generatePelletsMap(c.Pellets)

	for _, pelletFlow := range c.PelletsFlows {
		//sort flow order asc
		sort.SliceStable(pelletFlow.Flow, func(i, j int) bool {
			return pelletFlow.Flow[i].OrderID < pelletFlow.Flow[j].OrderID
		})
		asyncPelletFlow := pelletFlow
		if pelletFlow.IsActive {
			go c.ValidatePelletFlow(&asyncPelletFlow, pelletsMap)
		}
	}
	return nil
}

func (c *ContractValidator) ValidatePelletFlow(pelletFlow *models.PelletFlow, pelletsMap map[int]models.Pellet) {
	// iteration across orders in flow
	for _, order := range pelletFlow.Flow {
		pellet, exist := pelletsMap[order.PelletID]
		if !exist {
			c.notifyError(fmt.Sprintf("pellet with ID:[%i] doesn't exists - ValidatePelletFlow", order.PelletID))
			continue
		}
		if pellet.IsActive {
			//log.Printf("running validator for pellet: %v and flow: %v ", pellet, pelletFlow.Name)
			err := c.validateAPIRequest(&pellet)
			if err != nil {
				c.notifyError(err.Error())
				continue
			}

			//success
			c.notifySuccess(fmt.Sprintf("pellet request validated successfully - pellet-id:%d pellet-flow-id:%d flow-order:%d flow-name:%s",
				pellet.ID, pelletFlow.ID, order.OrderID, pelletFlow.Name))
		}
	}
}

func (c *ContractValidator) validateAPIRequest(pellet *models.Pellet) error {
	requestBody, err := json.Marshal(pellet.Request.Data)
	if err != nil {
		return fmt.Errorf("An error ocurred trying to parse body for ValidateAPIRequest:%s", err)
	}

	req, err := http.NewRequest(string(pellet.HttpMethod), pellet.Url, bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("An error ocurred trying to create a http request for ValidateAPIRequest:%s", err)
	}

	for headerKey, headerValue := range *pellet.Request.Headers {
		req.Header.Set(headerKey, headerValue)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("An error ocurred trying to make http request for ValidateAPIRequest:%s", err)
	}
	defer resp.Body.Close()

	err = CompareExpectedResponseBody(pellet, resp)
	if err != nil {
		return err
	}
	return nil
}

func CompareExpectedResponseBody(pellet *models.Pellet, httpResponse *http.Response) error {
	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return fmt.Errorf("An error ocurred trying to create a http request for CompareExpectedResponseBody:%s", err)
	}

	//validate headers
	for expectedHeaderKey, expectedHeaderValue := range *pellet.Response.Headers {
		responseHeader := httpResponse.Header.Get(strings.ToLower(expectedHeaderKey))
		if strings.ToLower(responseHeader) != strings.ToLower(expectedHeaderValue) {
			return fmt.Errorf("header:[%s] with expected value:[%s] is different from actual response header value:[%s] in CompareExpectedResponseBody:%s",
				expectedHeaderKey, expectedHeaderValue, responseHeader, err)
		}
	}

	//validate response code
	if httpResponse.StatusCode != pellet.Response.ResponseCode {
		return fmt.Errorf("expected response status code:%d is different from actual response statuscode:%d CompareExpectedResponseBody:%s",
			pellet.Response.ResponseCode, httpResponse.StatusCode, err)
	}

	//validate body
	response := map[string]interface{}{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("An error ocurred trying to parse http request body for CompareExpectedResponseBody:%s", err)
	}

	if !reflect.DeepEqual(*pellet.Response.Data, response) {
		return fmt.Errorf("expected response data is different from actual response data CompareExpectedResponseBody - actual: %v | expected: %v", response, pellet.Response.Data)
	}
	return nil
}

func (c *ContractValidator) notifyError(message string) {
	//utils.NotifyToSlack(message) //TODO: enabled this until we get the approval for nightwatcher-owl app in slack
	log.Println("[ERROR] -> " + message)
}

func (c *ContractValidator) notifySuccess(message string) {
	//utils.NotifyToSlack(message) //TODO: enabled this until we get the approval for nightwatcher-owl app in slack
	log.Println("[SUCCESS] -> " + message)
}

func (c *ContractValidator) generatePelletsMap(pellets []models.Pellet) map[int]models.Pellet {
	mapPellet := make(map[int]models.Pellet)

	for _, pellet := range pellets {
		mapPellet[pellet.ID] = pellet
	}

	return mapPellet
}
