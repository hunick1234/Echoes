package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Sender_Mail     string
	Sender_Password string
	RegisterSubject = "this is your login URL"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Sender_Mail = os.Getenv("sender_mail")
	Sender_Password = os.Getenv("sender_password")
	RegisterSubject = "this is your login URL"
}
