package main

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

type SessionHandler func(http.ResponseWriter, *http.Request, *sessions.Session)

type SessionValidator struct {
	handler  SessionHandler
	redirect string
}

func ValidateSession(handler SessionHandler, path string) SessionValidator {
	return SessionValidator{handler, path}
}

func (s SessionValidator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login")
	if err != nil {
		log.Println("[session validator] error reading session")
		http.Redirect(w, r, s.redirect, http.StatusFound)
		return
	}

	if session.IsNew {
		http.Redirect(w, r, s.redirect, http.StatusFound)
		return
	}
	s.handler(w, r, session)
}
