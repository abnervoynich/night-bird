package utils

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestJsonFileLoader(t *testing.T) {
	tempFilePath := "../../json-pellets/pellets/test.json"
	expectedJson := `{
  "name": "linkedin-example",
  "url": "www.-.com",
  "request": {
    "headers": {
      "X-RestLi-Protocol-Version": "2.0.0",
      "LinkedIn-Version": "202303",
      "authorization": "Bearer 12345"
    },
    "data": {}
  },
  "response": {
    "headers": {
      "X-RestLi-Protocol-Version": "2.0.0",
      "LinkedIn-Version": "202303",
      "authorization": "Bearer 12345"
    },
    "data": {},
    "response-code": 200
  }
}
`
	response, err := jsonFileLoader(tempFilePath)
	if err != nil {
		t.Fatalf("error trying to run JsonFileLoader: %s", err)
	}

	expectedMap := map[string]interface{}{}
	responseMap := map[string]interface{}{}

	err = json.Unmarshal(response, &responseMap)
	if err != nil {
		t.Errorf("error trying to json unmarshall response: %s", err)
	}

	err = json.Unmarshal([]byte(expectedJson), &expectedMap)
	if err != nil {
		t.Fatalf("error trying to json unmarshall expected response: %s", err)
	}

	if !reflect.DeepEqual(responseMap, expectedMap) {
		t.Fatalf("expected response is not correct")
	}
}

func TestGetLocalPellets(t *testing.T) {
	tempFilePath := "test.json"
	expectedJson := `{
  "name": "linkedin-example",
  "url": "www.-.com",
  "request": {
    "headers": {
      "X-RestLi-Protocol-Version": "2.0.0",
      "LinkedIn-Version": "202303",
      "authorization": "Bearer 12345"
    },
    "data": {}
  },
  "response": {
    "headers": {
      "X-RestLi-Protocol-Version": "2.0.0",
      "LinkedIn-Version": "202303",
      "authorization": "Bearer 12345"
    },
    "data": {},
    "response-code": 200
  }
}
`
	expectedMap := map[string][]byte{
		tempFilePath: []byte(expectedJson),
	}

	responseMap, err := GetLocalPellets()
	if err != nil {
		t.Fatalf("error trying to get local pellets: %s", err)
	}

	if !reflect.DeepEqual(responseMap, expectedMap) {
		t.Fatalf("expected response is not correct")
	}
}

func TestGetLocalPelletsFlows(t *testing.T) {
	tempFilePath := "test.json" //TODO: fix json paths problem (injectable directory)
	expectedJson := `{
  "id": 1,
  "name": "multiple requests for single case",
  "project-name": "service-social-communication",
  "pellets": [
    {
      "order-id": 1,
      "pellet-id": 1
    },
    {
      "order-id": 2,
      "pellet-id": 1
    }
  ]
}
`
	expectedMap := map[string][]byte{
		tempFilePath: []byte(expectedJson),
	}

	responseMap, err := GetLocalPelletsFlows()
	if err != nil {
		t.Fatalf("error trying to get local pellets flows: %s", err)
	}

	if !reflect.DeepEqual(responseMap, expectedMap) {
		t.Fatalf("expected response is not correct")
	}
}
