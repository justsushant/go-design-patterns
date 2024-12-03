package main

import (
	"encoding/json"
	"fmt"
)

func main() {

}

func ParseOdometer(input string) (int, error) {
	// unmarhalling into envelope type to determine the type of data
	m := ModelEnvelope{}
	if err := json.Unmarshal([]byte(input), &m); err != nil {
		return 0, fmt.Errorf("Error occured while unmarshaling into envelope type: %v", err)
	}

	// unmarshalling into specific type according to model
	switch m.Model {
	case "A":
		var data ModelAData
		if err := json.Unmarshal(m.Data, &data); err != nil {
			return 0, fmt.Errorf("Error occured while unmarshaling into ModelA type: %v", err)
		}
		return data.Odometer, nil

	default:
		return 0, nil
	}
}
