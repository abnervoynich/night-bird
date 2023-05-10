package utils

import (
	"encoding/json"
	"fmt"
	"github.com/abnervoynich/night-bird/app/models"
)

func ParseBytePellets(bytePellets *map[string][]byte) ([]models.Pellet, error) {
	pellets := []models.Pellet{}
	for fileName, pellet := range *bytePellets {
		pelletObj := models.Pellet{}
		err := json.Unmarshal(pellet, &pelletObj)
		if err != nil {
			return nil, fmt.Errorf("cannot unmarshall pellet from file: %s, error: %s", fileName, err)
		}

		pellets = append(pellets, pelletObj)
	}
	return pellets, nil
}

func ParseBytePelletsFlows(bytePelletsFlow *map[string][]byte) ([]models.PelletFlow, error) {
	pelletsFlows := []models.PelletFlow{}
	for fileName, pelletFlow := range *bytePelletsFlow {
		pelletFlowObj := models.PelletFlow{}
		err := json.Unmarshal(pelletFlow, &pelletFlowObj)
		if err != nil {
			return nil, fmt.Errorf("cannot unmarshall pellet flow from file: %s, error: %s", fileName, err)
		}

		pelletsFlows = append(pelletsFlows, pelletFlowObj)
	}
	return pelletsFlows, nil
}
