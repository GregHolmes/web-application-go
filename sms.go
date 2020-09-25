package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/vonage/vonage-go-sdk"
)

type SmsDetails struct {
	To      string
	Content string
}

type ReceivedSmsDetails struct {
	Msisdn           string
	To               string
	MessageId        string
	Text             string
	Type             string
	Keyword          string
	MessageTimestamp string
	ApiKey           string
}

func sendSms(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("src/template/send-sms.html")

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)

		return
	}

	smsDetails := SmsDetails{
		To:      r.FormValue("to"),
		Content: r.FormValue("content"),
	}

	auth := vonage.CreateAuthFromKeySecret(os.Getenv("VONAGE_API_KEY"), os.Getenv("VONAGE_API_SECRET"))
	smsClient := vonage.NewSMSClient(auth)
	response, err := smsClient.Send(os.Getenv("VONAGE_BRAND"), smsDetails.To, smsDetails.Content, vonage.SMSOpts{})

	if err != nil {
		panic(err)
	}

	if response.Messages[0].Status == "0" {
		fmt.Println("Account Balance: " + response.Messages[0].RemainingBalance)
	}

	tmpl.Execute(w, nil)
}

func receiveSms(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received an sms")

	// return
	receivedSmsDetails := ReceivedSmsDetails{
		Msisdn:           r.FormValue("msisdn"),
		To:               r.FormValue("to"),
		MessageId:        r.FormValue("messageId"),
		Text:             r.FormValue("text"),
		Type:             r.FormValue("type"),
		Keyword:          r.FormValue("keyword"),
		MessageTimestamp: r.FormValue("message-timestamp"),
		ApiKey:           r.FormValue("api-key"),
	}

	if receivedSmsDetails.ApiKey != os.Getenv("VONAGE_API_KEY") {
		fmt.Println("APIKEY doesn't match")
		return
	}

	message := Conversation{Msisdn: receivedSmsDetails.Msisdn, To: receivedSmsDetails.To, MessageId: receivedSmsDetails.MessageId, Text: receivedSmsDetails.Text, Type: receivedSmsDetails.Type, Keyword: receivedSmsDetails.Keyword, MessageTimestamp: receivedSmsDetails.MessageTimestamp}

	_ = db.Create(&message)

	fmt.Println("Message saved... check the database")

	return
}
