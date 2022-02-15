package middlewares

import (
	"encoding/json"
	"net/http"
	"sender/models"
)

func ErrorResponse(error string, writer http.ResponseWriter) {
	resp := &models.RespMsg{Statuscode: 400, Message: error}
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(writer).Encode(resp)
}

func ServerErrResponse(error string, writer http.ResponseWriter) {
	resp := &models.RespMsg{Statuscode: 500, Message: error}
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(writer).Encode(resp)
}

func SuccessResponse(msg string, writer http.ResponseWriter) {
	resp := &models.RespMsg{Statuscode: 200, Message: msg}
	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(resp)
}
