package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"sender/routes"
	"strings"
	"testing"
)

var route = routes.Routes()

func TestSendRequestEndpoint(t *testing.T) {
	request := url.Values{
		"name":    []string{"Jhon Doe"},
		"phone":   []string{"99999999"},
		"subject": []string{"Subject message"},
		"message": []string{"Test message"},
	}

	req, err := http.NewRequest("POST", "/api/v1/send/request", strings.NewReader(request.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), "success") {
		t.Errorf("handler returned wrong status code: got %v", rr.Body.String())
	}
}
