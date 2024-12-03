package main

import (
	"errors"
	"testing"
)

func TestMain(t *testing.T) {
	tt := []struct {
		name   string
		input  string
		expOut int
		expErr error
	}{
		{
			name: "model A json with odometer",
			input: `{
				"model": "A",
				"data": {
					"odometer": 244433,
					"fuelLevel": 55
				}
			}`,
			expOut: 244433,
			expErr: nil,
		},
		{
			name: "model A json with odometer but zero",
			input: `{
				"model": "A",
				"data": {
					"odometer": 0,
					"fuelLevel": 55
				}
			}`,
			expOut: 0,
			expErr: nil,
		},
		{
			name: "model A json without odometer",
			input: `{
				"model": "A",
				"data": {
					"fuelLevel": 55
				}
			}`,
			expErr: ErrOdometerNotFound,
		},
		{
			name: "model B json with odometer",
			input: `{
				"model": "B",
				"data": {
					"fuel": {
						"fuelLevel": 18
					},
					"speed": {
						"odometer": 51200
					}
				}
			}`,
			expOut: 51200,
			expErr: nil,
		},
		{
			name: "model B json with odometer but zero",
			input: `{
				"model": "B",
				"data": {
					"fuel": {
						"fuelLevel": 18
					},
					"speed": {
						"odometer": 0
					}
				}
			}`,
			expOut: 0,
			expErr: nil,
		},
		{
			name: "model B json without odometer",
			input: `{
				"model": "B",
				"data": {
					"fuel": {
						"fuelLevel": 18
					}
				}
			}`,
			expOut: 512,
			expErr: ErrOdometerNotFound,
		},
		{
			name: "model C json with odometer",
			input: `{
				"model": "C",
				"data": [
					{
						"spnId": 147,
						"pgnId": "F023",
						"spnName": "fuelLevel",
						"value": 35
					},
					{
						"spnId": 102,
						"pgnId": "F003",
						"spnName": "odometer",
						"value": 6349303
					}
				]
			}`,
			expOut: 6349303,
			expErr: nil,
		},
		{
			name: "model C json with odometer but zero",
			input: `{
				"model": "C",
				"data": [
					{
						"spnId": 147,
						"pgnId": "F023",
						"spnName": "fuelLevel",
						"value": 35
					},
					{
						"spnId": 102,
						"pgnId": "F003",
						"spnName": "odometer",
						"value": 0
					}
				]
			}`,
			expOut: 0,
			expErr: nil,
		},
		{
			name: "model C json without odometer",
			input: `{
				"model": "C",
				"data": [
					{
						"spnId": 147,
						"pgnId": "F023",
						"spnName": "fuelLevel",
						"value": 35
					}
				]
			}`,
			expErr: ErrOdometerNotFound,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			od, err := ParseOdometer(tc.input)
			if tc.expErr != nil {
				if err == nil {
					t.Fatalf("Expected error %q but got nil", tc.expErr.Error())
				}

				if !errors.Is(err, tc.expErr) {
					t.Errorf("Expected error %q but got %q", tc.expErr.Error(), err.Error())
				}

				return
			}
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if od != tc.expOut {
				t.Errorf("Expected %d but got %d", tc.expOut, od)
			}
		})
	}
}
