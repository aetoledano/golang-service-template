package models

type ApiMessage struct {
	Code int         `json:"code"`
	Tag  string      `json:"tag"`
	Data interface{} `json:"data,omitempty"`
}

type BusinessError struct {
	Code int    `json:"code"`
	Tag  string `json:"tag"`
}
