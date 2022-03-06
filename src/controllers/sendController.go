package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"reflect"
	"sender/middlewares"
	"sender/models"
	"strings"
)

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
var SendRequest = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var req models.Request
	reqBody, _ := ioutil.ReadAll(request.Body)
	request.Body.Close()
	request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		if err = request.ParseForm(); err != nil {
			middlewares.ServerErrResponse(err.Error(), response)
			return
		} else {
			pointer := reflect.ValueOf(&req)
			fields := pointer.Elem()
			if fields.Kind() == reflect.Struct {
				for k, v := range request.PostForm {
					k = strings.Title(strings.ToLower(k))
					field := fields.FieldByName(k)
					if field.IsValid() {
						if field.CanSet() {
							if field.Kind() == reflect.String {
								field.SetString(v[0])
							}
						}
					}
				}
			}
		}
	}
	host := middlewares.DotEnvVariable("SMTP_HOST")
	login := middlewares.DotEnvVariable("SMTP_LOGIN")
	password := middlewares.DotEnvVariable("SMTP_PASSWORD")
	port := middlewares.DotEnvVariable("SMTP_PORT")

	subject := middlewares.DotEnvVariable("SMTP_SUBJECT")
	message := middlewares.DotEnvVariable("SMTP_MESSAGE")

	unpackReq := reflect.ValueOf(req)
	unpackKeys := unpackReq.Type()

	for i := 0; i < unpackReq.NumField(); i++ {
		message = strings.ReplaceAll(message, "#"+strings.ToLower(unpackKeys.Field(i).Name)+"#", unpackReq.Field(i).String())
	}

	message = "Subject:" + subject + "\r\n\r\n" + message

	to := []string{middlewares.DotEnvVariable("SMTP_TO")}

	body := []byte(message)
	auth := smtp.PlainAuth("", login, password, host)

	err = smtp.SendMail(host+":"+port, auth, login, to, body)
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	middlewares.SuccessResponse("Send success", response)
})
