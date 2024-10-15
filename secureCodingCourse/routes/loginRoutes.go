package routes

import (
	"net/http"
	"secureCodingCourse/controller"
	"secureCodingCourse/db"
)

func SetUpSingleFactorAuthRoutes(c controller.IController, db map[string]db.IDB) {
	http.HandleFunc("/singlefactorlogin", func(w http.ResponseWriter, r *http.Request) {
		c.LoginUserSingleFactor(w, r, db)
	})
}
