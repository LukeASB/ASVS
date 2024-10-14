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
	var databases = make(map[string]db.IDB)

	c := controller.NewController()
	globomanticsDB, err := db.NewDB("Globomantics")

	if err != nil {
		log.Fatal("💣")
	}

	databases["Globomantics"] = globomanticsDB

	defer cleanup(databases)

	routes.SetUpRoutes(c, databases)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("💣")
	}
}

func cleanup(db map[string]db.IDB) {
	for _, val := range db {
		val.Close()
	}
}
