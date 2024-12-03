package main

import (
	"encoding/json"
	"fmt"
)

type ModelEnvelope struct {
	Model string          `json:"model"`
	Data  json.RawMessage `json:"data"`
}

type ModelA struct {
	Model string     `json:"model"`
	Data  ModelAData `json:"data"`
}

type ModelAData struct {
	Odometer int `json:"odometer"`
}

func main() {

}

func ParseOdometer(input string) (int, error) {
	m := ModelEnvelope{}
	if err := json.Unmarshal([]byte(input), &m); err != nil {
		return 0, fmt.Errorf("Error occured while unmarshaling into envelope type: %v", err)
	}

	if m.Model == "A" {
		var data ModelAData
		if err := json.Unmarshal(m.Data, &data); err != nil {
			return 0, fmt.Errorf("Error occured while unmarshaling into ModelA type: %v", err)
		}

		return data.Odometer, nil
	}

	return 0, nil
}
