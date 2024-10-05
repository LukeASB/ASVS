package controller

import (
	"fmt"
	"net/http"
	"secureCodingCourse/validation"
)

type Controller struct{}

type IController interface {
	Whitelist(w http.ResponseWriter, r *http.Request)
	BoundaryCheck(w http.ResponseWriter, r *http.Request)
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

func (c *Controller) BoundaryCheck(w http.ResponseWriter, r *http.Request) {
	inputNum1 := r.URL.Query().Get("num1")
	inputNum2 := r.URL.Query().Get("num2")

	num1, err := validation.NewBoundaryCheck().Validate(inputNum1)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	num2, err := validation.NewBoundaryCheck().Validate(inputNum2)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("%d + %d = %d", num1, num2, num1+num2)))
}
