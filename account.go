package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func account(w http.ResponseWriter, r *http.Request) {
	isAuthenticated(w, r)

	tmpl, err := template.ParseFiles("src/template/account.html")

	if err != nil {
		fmt.Println(err)
	}

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
}
