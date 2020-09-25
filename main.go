package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type LoginDetails struct {
	Email    string
	Password string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectDb()

	mux := mux.NewRouter()

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/account", account)
	mux.HandleFunc("/register", register)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/send-sms", sendSms)
	mux.HandleFunc("/receive-sms", receiveSms)
	mux.HandleFunc("/verify-send", sendVerificationCode)
	mux.HandleFunc("/verify-code/{requestId}", verifyMyCode)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", mux)
}
