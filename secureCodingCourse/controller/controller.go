package controller

import (
	"fmt"
	"net/http"
	"secureCodingCourse/validation"
)

type Controller struct{}

type IController interface {
	Whitelist(w http.ResponseWriter, r *http.Request)
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Whitelist(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")

	valid := validation.NewWhiteList().Check(input)

	if valid {
		w.Write([]byte(fmt.Sprintf("%s is a good input!", input)))
		return
	}

	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
