package services_test

import (
	"encoding/json"
	"github.com/abnervoynich/night-bird/app/models"
	"github.com/abnervoynich/night-bird/app/services"
	"github.com/abnervoynich/night-bird/app/utils"
	"log"
	"reflect"
	"testing"
)

func TestLoadLocalPellets(t *testing.T) {
	mocker := services.NewMockerService()

	bytePellets, err := utils.GetJsonFromLocalPellets()
	if err != nil {
		t.Fatalf("cannot get byte pellets from utils.GetJsonFromLocalPellets: %s", err)
	}

	pellets := []models.Pellet{}
	for fileName, pellet := range bytePellets {
		pelletObj := models.Pellet{}
		err = json.Unmarshal(pellet, &pelletObj)
		if err != nil {
			log.Printf("cannot unmarshall pellet from file: %s, error: %s", fileName, err)
		}

		pellets = append(pellets, pelletObj)
	}

	if !reflect.DeepEqual(pellets, mocker.Pellets) {
		t.Fatalf("expected response is not as the current response")
	}
}
