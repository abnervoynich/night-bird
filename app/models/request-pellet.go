package models

type RequestPellet struct {
	Headers *map[string]string      `json:"headers"`
	Data    *map[string]interface{} `json:"data"`
}
