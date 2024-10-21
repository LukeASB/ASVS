package routes

import (
	"net/http"
	"secureCodingCourse/controller"
)

func Routes(c controller.IController) {
	http.HandleFunc("/hashDemo", func(w http.ResponseWriter, r *http.Request) {
		c.HashDemo(w, r)
	})

	http.HandleFunc("/aesencryptdecryptdemo", func(w http.ResponseWriter, r *http.Request) {
		c.AESEncryptDecryptDemo(w, r)
	})
}
