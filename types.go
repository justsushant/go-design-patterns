package main

import "encoding/json"

type ModelEnvelope struct {
	Model string          `json:"model"`
	Data  json.RawMessage `json:"data"`
}

type Model string

var (
	ModelA = Model("A")
	ModelB = Model("B")
)

type ModelAData struct {
	Odometer int `json:"odometer"`
}

type ModelBData struct {
	Speed ModelBSpeedData `json:"speed"`
}

type ModelBSpeedData struct {
	Odometer int `json:"odometer"`
}
