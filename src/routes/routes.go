package routes

import (
	"sender/controllers"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/send/request", controllers.SendRequest).Methods("POST")
	return router
}
