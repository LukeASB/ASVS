package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"secureCodingCourse/data"
	"secureCodingCourse/db"
	"secureCodingCourse/validation"

	_ "github.com/denisenkom/go-mssqldb" // MS SQL driver
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
	UnsecureCrossSiteScriptingExample(w http.ResponseWriter, r *http.Request)
	SecureCrossSiteScriptingExample(w http.ResponseWriter, r *http.Request)
	SQLInjection(w http.ResponseWriter, r *http.Request, db db.IDB)
	SafeSQLSearchExample(w http.ResponseWriter, r *http.Request, db db.IDB)
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Whitelist(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")

	if len(input) <= 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

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

	if len(inputNum1) <= 0 || len(inputNum2) <= 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

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
	response := map[string]bool{
		"IsNumber": false,
	}

	input := r.URL.Query().Get("value")

	if len(input) <= 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response["IsNumber"] = validation.NewNumericValidation().IsNumber(input)

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func (c *Controller) CheckForNullBytes(w http.ResponseWriter, r *http.Request) {
	response := map[string]bool{
		"HasNullByte": false,
	}

	input := r.URL.Query().Get("text")

	if len(input) <= 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response["HasNullByte"] = validation.NewCheckForNullBytes().ContainsNullBytes(input)

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func (c *Controller) NewLineCharacterEscape(w http.ResponseWriter, r *http.Request) {
	response := map[string]bool{
		"HasNewLineCharacterEscape": false,
	}

	input := r.URL.Query().Get("text")

	if len(input) <= 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response["HasNewLineCharacterEscape"] = validation.NewNewLineCharacterEscape().CheckForNewLineCharacter(input)

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func (c *Controller) PathAlterationEscape(w http.ResponseWriter, r *http.Request) {
	response := map[string]bool{
		"IsValidPath": false,
	}

	input := r.URL.Query().Get("input")

	if len(input) <= 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response["IsValidPath"] = validation.NewPathAlterationEscape().IsValidPath(input)

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func (c *Controller) CheckForExtendedUTF8Encoding(w http.ResponseWriter, r *http.Request) {
	response := map[string]bool{
		"IsValidUTF8": false,
	}

	input := r.URL.Query().Get("input")

	if len(input) <= 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response["IsValidUTF8"] = validation.NewCheckForExtendedUTF8Encoding().IsValidUTF8(input)

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func (c *Controller) UnsecureCrossSiteScriptingExample(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")

	if len(input) <= 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("<h1>Hi Cross Site Scripting Example! %s </h1>", input)))
}

func (c *Controller) SecureCrossSiteScriptingExample(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("input")

	if len(input) <= 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("output").Parse("<h1>Hello, {{ . }}!</h1>")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, input)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (c *Controller) SQLInjection(w http.ResponseWriter, r *http.Request, db db.IDB) {
	input := r.URL.Query().Get("input")

	if len(input) <= 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	results, err := db.UnsafeRetrieve(input)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	results, ok := results.([]data.Patient)

	if !ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Println("Patients data:")

	jsonData, err := json.Marshal(results)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}

func (c *Controller) SafeSQLSearchExample(w http.ResponseWriter, r *http.Request, db db.IDB) {
	input := r.URL.Query().Get("input")

	if len(input) <= 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	results, err := db.SafeRetrieve(input)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	results, ok := results.([]data.Patient)

	if !ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Println("Patients data:")

	jsonData, err := json.Marshal(results)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}
