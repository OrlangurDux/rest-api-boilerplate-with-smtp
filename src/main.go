package main

import (
	"log"
	"net/http"
	"sender/middlewares"
	"sender/routes"

	"github.com/rs/cors"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	port := middlewares.DotEnvVariable("PORT")

	router := routes.Routes()

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-type", "Origin", "Accept", "*"},
	})

	handler := c.Handler(router)

	http.ListenAndServe(":"+port, middlewares.LogRequest(handler))
}
