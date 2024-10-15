package routes

import (
	"net/http"
	"secureCodingCourse/controller"
	"secureCodingCourse/db"
)

func SetUpMultiFactorAuthRoutesDemo(c controller.IMultiFactorDemoController, db map[string]db.IDB) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c.RedirectToHomeHandler(w, r)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		c.LoginHandler(w, r)
	})

	http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		c.VerifyHandler(w, r)
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		c.HomeRedirectHandler(w, r)
	})

	http.HandleFunc("/home.html", func(w http.ResponseWriter, r *http.Request) {
		c.HomeHandler(w, r)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("ui/static"))))
}
