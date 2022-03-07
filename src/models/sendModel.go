package models

//RespMsg -> string message response
type RespMsg struct {
	Statuscode int    `json:"string"`
	Message    string `json:"msg"`
} //@name Response

//Request -> request message
type Request struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
} //@name Request
