package main

import (
	"log"
	"net/http"
	"secureCodingCourse/controller"
	"secureCodingCourse/db"
	"secureCodingCourse/routes"

	_ "github.com/denisenkom/go-mssqldb" // MS SQL driver
)

func main() {
	c := controller.NewController()
	db, err := db.NewDB()

	if err != nil {
		log.Fatal("💣")
	}

	defer cleanup(db)

	routes.SetUpRoutes(c, db)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("💣")
	}
}

func cleanup(db db.IDB) {
	db.Close()
}
