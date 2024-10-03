package routes

import (
	"net/http"
	"secureCodingCourse/controller"
)

func SetUpRoutes(c controller.IController) {
	http.HandleFunc("/whitelist", func(w http.ResponseWriter, r *http.Request) {
		c.Whitelist(w, r)
	})
}
