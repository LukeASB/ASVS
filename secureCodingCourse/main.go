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
	mFDc := controller.NewMultiFactorDemoController()
	globomanticsDB, err := db.NewDB("Globomantics")

	if err != nil {
		log.Fatal("💣")
	}

	databases["Globomantics"] = globomanticsDB

	defer cleanup(databases)

	routes.SetUpValidationRoutes(c, databases)
	routes.SetUpSingleFactorAuthRoutes(c, databases)
	routes.SetUpMultiFactorAuthRoutesDemo(mFDc, databases)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func cleanup(db map[string]db.IDB) {
	for _, val := range db {
		val.Close()
	}
}
