package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Email       string
	Password    string
	PhoneNumber string
}

type Conversation struct {
	gorm.Model
	Msisdn           string
	To               string
	MessageId        string
	Text             string
	Type             string
	Keyword          string
	MessageTimestamp string
}

func connectDb() {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Conversation{})
}
