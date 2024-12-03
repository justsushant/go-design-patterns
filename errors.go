package main

import "errors"

var (
	ErrOdometerNotFound = errors.New("odometer not found")
	ErrModelNotFound    = errors.New("model not found")
)
