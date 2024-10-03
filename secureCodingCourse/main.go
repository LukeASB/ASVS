package main

import (
	"log"
	"net/http"
	"secureCodingCourse/controller"
	"secureCodingCourse/routes"
)

func main() {
	c := controller.NewController()

	routes.SetUpRoutes(c)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("💣")
	}
}
