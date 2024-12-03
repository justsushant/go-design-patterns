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
	switch Model(m.Model) {
	case ModelA:
		var data ModelAData
		if err := json.Unmarshal(m.Data, &data); err != nil {
			return 0, fmt.Errorf("Error occured while unmarshaling into ModelA type: %v", err)
		}

		if data.Odometer != nil {
			return *data.Odometer, nil
		}
		return 0, ErrOdometerNotFound

	case ModelB:
		var data ModelBData
		if err := json.Unmarshal(m.Data, &data); err != nil {
			return 0, fmt.Errorf("Error occured while unmarshaling into ModelB type: %v", err)
		}

		if data.Speed != nil && data.Speed.Odometer != nil {
			return *data.Speed.Odometer, nil
		}
		return 0, ErrOdometerNotFound

	case ModelC:
		var data []ModelCDataPoint
		if err := json.Unmarshal(m.Data, &data); err != nil {
			return 0, fmt.Errorf("Error occured while unmarshaling into ModelC type: %v", err)
		}

		for _, d := range data {
			if d.SpnName == "odometer" {
				return d.Value, nil
			}
		}
		return 0, ErrOdometerNotFound

	default:
		return 0, nil
	}
}
