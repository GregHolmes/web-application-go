package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vonage/vonage-go-sdk"
)

func sendVerificationCode(w http.ResponseWriter, r *http.Request) {
	auth := vonage.CreateAuthFromKeySecret(os.Getenv("VONAGE_API_KEY"), os.Getenv("VONAGE_API_SECRET"))
	verifyClient := vonage.NewVerifyClient(auth)

	response, errResp, err := verifyClient.Request(os.Getenv("NUMBER"), "GoTest", vonage.VerifyOpts{CodeLength: 6, Lg: "es-es", WorkflowID: 6})

	if err != nil {
		fmt.Printf("%#v\n", err)
	} else if response.Status != "0" {
		fmt.Println("Error status " + errResp.Status + ": " + errResp.ErrorText)
	} else {
		fmt.Println("Request started: " + response.RequestId)

		http.Redirect(w, r, "/verify-code/"+response.RequestId, http.StatusSeeOther)
	}
}

func verifyMyCode(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("src/template/verify-code.html")

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)

		return
	}

	auth := vonage.CreateAuthFromKeySecret(os.Getenv("VONAGE_API_KEY"), os.Getenv("VONAGE_API_SECRET"))
	verifyClient := vonage.NewVerifyClient(auth)

	vars := mux.Vars(r)

	response, errResp, err := verifyClient.Check(vars["requestId"], r.FormValue("verification-code"))

	if err != nil {
		fmt.Printf("%#v\n", err)
	} else if response.Status != "0" {
		fmt.Println("Error status " + errResp.Status + ": " + errResp.ErrorText)
	} else {
		// all good
		fmt.Println("Request complete: " + response.RequestId)
	}
}
