package main

import "encoding/json"

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
