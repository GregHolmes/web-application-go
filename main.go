package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type LoginDetails struct {
	Email    string
	Password string
}

func main() {

	mux := mux.NewRouter()

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/account", account)
	mux.HandleFunc("/register", register)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
