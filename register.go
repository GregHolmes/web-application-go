package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func register(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("src/template/register.html")

	if err != nil {
		fmt.Println(err)
	}

	if r.Method != http.MethodPost {
		fmt.Println("Register Not method post")
		fmt.Println(r.Method)
		tmpl.Execute(w, nil)

		return
	}
}
