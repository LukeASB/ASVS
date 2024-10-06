package routes

import (
	"net/http"
	"secureCodingCourse/controller"
)

func SetUpRoutes(c controller.IController) {
	http.HandleFunc("/whitelist", func(w http.ResponseWriter, r *http.Request) {
		c.Whitelist(w, r)
	})

	http.HandleFunc("/boundaryCheck", func(w http.ResponseWriter, r *http.Request) {
		c.BoundaryCheck(w, r)
	})

	http.HandleFunc("/characterescape", func(w http.ResponseWriter, r *http.Request) {
		c.CharacterEscape(w, r)
	})

	http.HandleFunc("/numericvalidation", func(w http.ResponseWriter, r *http.Request) {
		c.NumericValidation(w, r)
	})

	http.HandleFunc("/checkfornullbytes", func(w http.ResponseWriter, r *http.Request) {
		c.CheckForNullBytes(w, r)
	})

	http.HandleFunc("/newlinecharacterescape", func(w http.ResponseWriter, r *http.Request) {
		c.NewLineCharacterEscape(w, r)
	})

	http.HandleFunc("/pathalterationescape", func(w http.ResponseWriter, r *http.Request) {
		c.PathAlterationEscape(w, r)
	})

	http.HandleFunc("/checkforextendedutf8encoding", func(w http.ResponseWriter, r *http.Request) {
		c.CheckForExtendedUTF8Encoding(w, r)
	})

	http.HandleFunc("/unsecurecrosssitescriptingexample", func(w http.ResponseWriter, r *http.Request) {
		c.UnsecureCrossSiteScriptingExample(w, r)
	})

	http.HandleFunc("/securecrosssitescriptingexample", func(w http.ResponseWriter, r *http.Request) {
		c.SecureCrossSiteScriptingExample(w, r)
	})
}
