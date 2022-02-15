package models

//RespMsg -> string message response
type RespMsg struct {
	Statuscode int    `json:"string"`
	Message    string `json:"msg"`
}

//Request -> request message
type Request struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}
