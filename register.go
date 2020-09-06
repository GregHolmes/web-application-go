package main

import (
	"fmt"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func register(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("src/template/register.html")

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

	hashedPassword, _ := HashPassword(details.Password)

	fmt.Println(hashedPassword)

	details.Password = hashedPassword

	fmt.Println(hashedPassword)

	user1 := User{Email: details.Email, Password: details.Password}

	fmt.Println(user1.Email)
	_ = db.Create(&user1)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
