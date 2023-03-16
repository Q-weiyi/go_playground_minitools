package helper

import (
	"encoding/json"
)

func MinifyJSON(jsonStr string) ([]byte, error) {
	var jsonData map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &jsonData)
	if err != nil {
		return nil, err
	}

	minifiedJSON, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}

	return minifiedJSON, nil
}
