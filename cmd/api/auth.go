package main

import "net/http"

func (app *Application) EmailPassSignupHandler(w http.ResponseWriter, r *http.Request) {
	// Implement email-password signup logic here.
}

func (app *Application) GoogleOauthSignupHandler(w http.ResponseWriter, r *http.Request) {
	// Implement Google OAuth signup logic here.
}

func (app *Application) FacebookOauthSignupHandler(w http.ResponseWriter, r *http.Request) {
	// Implement Facebook OAuth signup logic here.
}
