package main

import "net/http"

func (app *application) EmailPassSignupHandler(w http.ResponseWriter, r *http.Request) {
	// Implement email-password signup logic here.
}

func (app *application) GoogleOauthSignupHandler(w http.ResponseWriter, r *http.Request) {
	// Implement Google OAuth signup logic here.
}

func (app *application) FacebookOauthSignupHandler(w http.ResponseWriter, r *http.Request) {
	// Implement Facebook OAuth signup logic here.
}
