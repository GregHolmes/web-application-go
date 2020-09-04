package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func account(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("src/template/account.html")

	if err != nil {
		fmt.Println(err)
	}

	if r.Method != http.MethodPost {
		fmt.Println("Account Not method post")
		fmt.Println(r.Method)
		tmpl.Execute(w, nil)
		return
	}
}
