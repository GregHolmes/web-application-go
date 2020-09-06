package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "go-project-auth")

	tmpl, err := template.ParseFiles("src/template/login.html")

	if err != nil {
		fmt.Println(err)
	}

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)

		return
	}

	details := LoginDetails{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	var user User

	result := db.Where("email = ?", r.FormValue("email")).First(&user)

	if result.RowsAffected != 1 {
		// Throw error, no user found
		fmt.Println("No user found")

		return
	}

	hashPassResult := CheckPasswordHash(details.Password, user.Password)

	if hashPassResult == false {
		// Throw error, passwords don't match

		fmt.Println("Passwords don't match")

		return
	}

	// Set a session for this log in..
	session.Values["authenticated"] = true
	session.Save(r, w)

	http.Redirect(w, r, "/account", http.StatusSeeOther)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func isAuthenticated(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "go-project-auth")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/login", http.StatusForbidden)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "go-project-auth")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
