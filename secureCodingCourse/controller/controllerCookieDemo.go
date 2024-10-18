package controller

import (
	"net/http"
	"time"
)

type ControllerCookieDemo struct{}

type IControllerCookieDemo interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	LoggedInAccessiblePage(w http.ResponseWriter, r *http.Request)
}

func NewControllerCookieDemo() *ControllerCookieDemo {
	return &ControllerCookieDemo{}
}

func (c *ControllerCookieDemo) Login(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "LSBCookie",
		Value:    "cookieToken",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(5 * time.Minute),
	}

	http.SetCookie(w, cookie)

	w.Write([]byte("Logged In! Cookie Set"))
}

func (c *ControllerCookieDemo) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "LSBCookie",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Unix(0, 0),
	}

	http.SetCookie(w, cookie)

	w.Write([]byte("Logged out! Deleted Cookie"))
}

func (c *ControllerCookieDemo) LoggedInAccessiblePage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("LSBCookie")

	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if cookie.Value != "cookieToken" {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	w.Write([]byte("Welcome! Logged In"))
}
