package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"secureCodingCourse/validation"
)

type Controller struct{}

type IController interface {
	Whitelist(w http.ResponseWriter, r *http.Request)
	BoundaryCheck(w http.ResponseWriter, r *http.Request)
	CharacterEscape(w http.ResponseWriter, r *http.Request)
	NumericValidation(w http.ResponseWriter, r *http.Request)
	CheckForNullBytes(w http.ResponseWriter, r *http.Request)
	NewLineCharacterEscape(w http.ResponseWriter, r *http.Request)
	PathAlterationEscape(w http.ResponseWriter, r *http.Request)
	CheckForExtendedUTF8Encoding(w http.ResponseWriter, r *http.Request)
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

func (c *Controller) CharacterEscape(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("text")

	sanitisedText := validation.NewCharacterEscape().EscapeHTML(input)

	w.Write([]byte(sanitisedText))
}

func (c *Controller) NumericValidation(w http.ResponseWriter, r *http.Request) {
	jsonResponse := map[string]bool{
		"IsNumber": false,
	}

	input := r.URL.Query().Get("value")

	jsonResponse["IsNumber"] = validation.NewNumericValidation().IsNumber(input)

	response, err := json.Marshal(jsonResponse)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(response)
}

func (c *Controller) CheckForNullBytes(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) NewLineCharacterEscape(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) PathAlterationEscape(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) CheckForExtendedUTF8Encoding(w http.ResponseWriter, r *http.Request) {

}
