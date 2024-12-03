package main

import "testing"

func TestMain(t *testing.T) {
	tt := []struct {
		name   string
		input  string
		expOut int
	}{
		{
			name: "model A json",
			input: `{
				"model": "A",
				"data": {
					"odometer": 243
				}
			}`,
			expOut: 243,
		},
		{
			name: "model B json",
			input: `{
				"model": "B",
				"data": {
					"speed": {
						"odometer": 512
					}
				}
			}`,
			expOut: 512,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			od, err := ParseOdometer(tc.input)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if od != tc.expOut {
				t.Errorf("Expected %d but got %d", tc.expOut, od)
			}
		})
	}
}
