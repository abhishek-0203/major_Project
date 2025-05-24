package projects

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func LoadProjectsFromFile(filePath string) ([]Project, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	// Unmarshalling the JSON data into a struct
	var wrapper struct { // struct wrapper
		Projects []Project `json:"projects"`
	}

	err = json.Unmarshal(byteValue, &wrapper)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %v", err)
	}

	return wrapper.Projects, nil
}
