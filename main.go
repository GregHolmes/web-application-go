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
	connectDb()

	mux := mux.NewRouter()

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/account", account)
	mux.HandleFunc("/register", register)
	mux.HandleFunc("/logout", logout)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
