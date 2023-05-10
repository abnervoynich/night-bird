package models

type ResponsePellet struct {
	Headers      *map[string]string      `json:"headers"`
	Data         *map[string]interface{} `json:"data"`
	ResponseCode int                     `json:"response-code"`
}
