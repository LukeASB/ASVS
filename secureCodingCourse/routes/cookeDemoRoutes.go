package routes

import (
	"net/http"
	"secureCodingCourse/controller"
)

func SetUpCookieDemoRoutes(c controller.IControllerCookieDemo) {
	http.HandleFunc("/v1/cookiedemo/Login", func(w http.ResponseWriter, r *http.Request) {
		c.Login(w, r)
	})

	http.HandleFunc("/v1/cookiedemo/Logout", func(w http.ResponseWriter, r *http.Request) {
		c.Logout(w, r)
	})

	http.HandleFunc("/v1/cookiedemo/LoggedInPage", func(w http.ResponseWriter, r *http.Request) {
		c.LoggedInAccessiblePage(w, r)
	})
}
