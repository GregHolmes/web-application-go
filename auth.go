package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
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

	// TODO: Compare user to any details in the database
	_ = details

	http.Redirect(w, r, "/account", http.StatusSeeOther)
}

// func setSession(userName string, response http.ResponseWriter) {
// 	value := map[string]string{
// 		"name": userName,
// 	}
// 	if encoded, err := cookieHandler.Encode("session", value); err == nil {
// 		cookie := &http.Cookie{
// 			Name:  "session",
// 			Value: encoded,
// 			Path:  "/",
// 		}
// 		http.SetCookie(response, cookie)
// 	}
// }

func clearSession() {

}

func logout() {

}
