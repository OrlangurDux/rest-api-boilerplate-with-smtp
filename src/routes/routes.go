package routes

import (
	"sender/controllers"
	_ "sender/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title        RESTful API
// @version      1.0
// @description  This is a boilerplate backend server resource.

// @contact.name  API Support

// @host      localhost:9769
// @BasePath  /api/v1

func Routes() *mux.Router {
	router := mux.NewRouter()
	c := controllers.BaseController()
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/send/request", c.SendRequest).Methods("POST")

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	return router
}
