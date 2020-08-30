package models

type ApiMessage struct {
	Code    int         `json:"code"`
	Subcode string      `json:"subCode,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
