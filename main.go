package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-not-very-secret"))
var templates = template.Must(template.ParseGlob("templates/*.html"))

func betHandler(w http.ResponseWriter, r *http.Request, s *sessions.Session) {
	user := s.Values["username"]
	choice := r.FormValue("choice")
	amount := r.FormValue("amount")

	log.Printf("%v bet %v on %v", user, amount, choice)

	http.Redirect(w, r, "/", http.StatusFound)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login")
	if err == nil {
		session.Values["username"] = r.FormValue("username")
		session.Save(r, w)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func logoutHandler(w http.ResponseWriter, r *http.Request, s *sessions.Session) {
	s.Options.MaxAge = -1
	s.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login")
	if err != nil {
		log.Println("error reading session")
	}
	view := templates.Lookup("main.html")
	if session.IsNew {
		view = templates.Lookup("login.html")
	}
	view.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/login", loginHandler)
	http.Handle("/logout", ValidateSession(logoutHandler, "/"))
	http.Handle("/bet", ValidateSession(betHandler, "/"))
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
