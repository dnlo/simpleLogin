package handlers

import (
	"html/template"
	"log"
	"net/http"
	"fmt"
	"github.com/dnlo/web/simpleLogin/db"
	"time"
)

func InitHandlers() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/internal", internalHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/admin", adminHandler)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	err := db.AuthUser(r.FormValue("name"), r.FormValue("password"))

	if err != nil {
		log.Printf("Failed to login: %v", err)
		http.Redirect(w, r, "/", 302)
		return
	}
	user, err := db.GetUser(r.FormValue("name"))
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "username", Value: user.Name, Expires: expiration}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/internal", 302)
}

func internalHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		log.Printf("Failed to open internal page: %v", err)
		http.Redirect(w, r, "/", 302)
		return
	}
	tmpl := template.Must(template.ParseFiles("./web/internal.html"))
	tmpl.Execute(w, cookie.Value)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		log.Printf("Failed to open internal page: %v", err)
		http.Redirect(w, r, "/", 302)
		return
	}
	user, err := db.GetUser(cookie.Value)
	if !user.Admin {
		fmt.Fprint(w, "You are not authorized to access the admin page")
		return
	}
	
	users := db.GetUserList()
	tmpl := template.Must(template.ParseFiles("./web/admin.html"))
	tmpl.Execute(w, &users)
}


func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "username",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 302)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	user, err := db.CreateUser(r.FormValue("name"), r.FormValue("password"), true)
	if err != nil {
		log.Printf("Failed to create new user: %v", err)
		http.Redirect(w, r, "/", 302)
		return
	}
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "username", Value: user.Name, Expires: expiration}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/internal", 302)
}
