package utils

import (
	"fmt"
	"io"
	"os"
)

const (
	pelletsDirectory      string = "./json-pellets/pellets/"
	pelletsFlowsDirectory string = "./json-pellets/pellets-flows/"
)

func jsonFileLoader(filepath string) ([]byte, error) {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	byteFile, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	return byteFile, nil
}

func getJsonFromLocalPellets(pelletsDirectory string) (*map[string][]byte, error) {
	jsonObjects := map[string][]byte{}
	entries, err := os.ReadDir(pelletsDirectory)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			fmt.Printf("an error occurred trying to get info from file: %s", err)
		}
		jsonObjects[info.Name()], err = jsonFileLoader(pelletsDirectory + info.Name())
	}

	return &jsonObjects, nil
}

func GetLocalPellets() (*map[string][]byte, error) {
	return getJsonFromLocalPellets(pelletsDirectory)
}

func GetLocalPelletsFlows() (*map[string][]byte, error) {
	return getJsonFromLocalPellets(pelletsFlowsDirectory)
}
