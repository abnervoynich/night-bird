package services

import (
	"encoding/json"
	"fmt"
	"github.com/abnervoynich/night-bird/app/models"
	"net/http"
	"reflect"
	"strings"
)

type Mocker struct {
	Pellets []models.Pellet
}

func NewMockerService(pellets []models.Pellet) *Mocker {
	return &Mocker{Pellets: pellets}
}

func (m *Mocker) ServeMockResponse(request *http.Request) (*models.ResponsePellet, error) {
	url := strings.Split(request.URL.RawQuery, "=")[1]
	data := map[string]interface{}{}
	method := request.Method

	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("cannot decode body request: %s", err)
	}

	for _, pellet := range m.Pellets {
		if pellet.Url == url && method == string(pellet.HttpMethod) {
			for headerName, headerVal := range *pellet.Request.Headers {
				requestHeaderVal := request.Header.Get(headerName)

				if requestHeaderVal == "" {
					return nil, fmt.Errorf("the request for url: %s doesn't contain header: %s as the pellet.response.header", url, headerName)
				}

				if strings.ToLower(requestHeaderVal) != strings.ToLower(headerVal) {
					return nil, fmt.Errorf("the request header: %s value:%s for url: %s isn't the same as the pellet.response.header value: %s", headerName, requestHeaderVal, url, headerVal)
				}
			}

			if !reflect.DeepEqual(data, *pellet.Request.Data) {
				return nil, fmt.Errorf("the request data: %v for url: %s isn't the same as the pellet.response data: %v", data, url, pellet.Request.Data)
			}

			return pellet.Response, nil
		}
	}

	return nil, fmt.Errorf("cannot find an specific pellet for the selected url")
}
