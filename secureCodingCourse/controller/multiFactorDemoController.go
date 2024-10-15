package controller

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image/png"
	"log"
	"net/http"
	"secureCodingCourse/data"
	"time"

	"github.com/pquerna/otp/totp"
)

var userSecretKeys = make(map[string]string)

type MultiFactorDemoController struct{}

type IMultiFactorDemoController interface {
	RedirectToHomeHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
	VerifyHandler(w http.ResponseWriter, r *http.Request)
	HomeRedirectHandler(w http.ResponseWriter, r *http.Request)
	HomeHandler(w http.ResponseWriter, r *http.Request)
}

func NewMultiFactorDemoController() *MultiFactorDemoController {
	return &MultiFactorDemoController{}
}

func (c *MultiFactorDemoController) RedirectToHomeHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home.html", http.StatusSeeOther)
}

func (c *MultiFactorDemoController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")

		// Check if the user already has a secret key stored
		secretKey, ok := userSecretKeys[username]
		if !ok {
			// Generate TOTP key
			key, err := totp.Generate(totp.GenerateOpts{
				Issuer:      "globomantics.com",
				AccountName: username,
			})
			if err != nil {
				log.Fatal(err)
			}

			// Store the secret key
			userSecretKeys[username] = key.Secret()
			secretKey = key.Secret()

			// Convert TOTP key into a PNG image
			img, err := key.Image(200, 200)
			if err != nil {
				log.Fatal(err)
			}

			// Encode PNG image to byte slice
			var buf bytes.Buffer
			err = png.Encode(&buf, img)
			if err != nil {
				log.Fatal(err)
			}

			qrCodeImage := base64.StdEncoding.EncodeToString(buf.Bytes())

			// Prepare the data for the template
			data := data.TemplateData{
				Username:    username,
				QRCodeImage: qrCodeImage,
				SecretKey:   secretKey,
			}

			tmpl, err := template.ParseFiles("ui/html/verify.html")
			if err != nil {
				log.Fatal(err)
			}

			err = tmpl.Execute(w, data)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			// User already has a secret key, display the verification form directly
			data := data.TemplateData{
				Username:  username,
				SecretKey: secretKey,
			}

			tmpl, err := template.ParseFiles("ui/html/verify.html")
			if err != nil {
				log.Fatal(err)
			}

			err = tmpl.Execute(w, data)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		// Render the login page
		tmpl, err := template.ParseFiles("ui/html/login.html")
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (c *MultiFactorDemoController) VerifyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Retrieve the passcode and secret key from the form
		passcode := r.FormValue("passcode")
		secretKey := r.FormValue("secretKey")
		username := r.FormValue("username")

		// Validate the TOTP passcode
		valid := totp.Validate(passcode, secretKey)
		if valid {
			// Set the login cookie
			expiration := time.Now().UTC().Add(24 * time.Hour) // Set the cookie expiration time (e.g., 24 hours)
			cookie := http.Cookie{
				Name:     "user",
				Value:    username,
				Expires:  expiration,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)

			// Redirect to the home page
			http.Redirect(w, r, "/home.html", http.StatusSeeOther)
			return
		} else {
			// Authentication failed, show error message
			fmt.Fprintf(w, "Authentication failed!")
			return
		}
	}

	// Handle GET request for the verification page
	tmpl, err := template.ParseFiles("ui/html/verify.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *MultiFactorDemoController) HomeRedirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home.html", http.StatusSeeOther)
}

func (c *MultiFactorDemoController) HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the login cookie exists
	cookie, err := r.Cookie("user")
	if err != nil || cookie.Value == "" {
		// Cookie does not exist or is empty, redirect to the login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	http.ServeFile(w, r, "ui/html/home.html")
}
