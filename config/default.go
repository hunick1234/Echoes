package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Sender_Mail     string
	Sender_Password string
	DB_Password     string
	DB_Name         string
	DB_User         string
	RegisterSubject = "this is your login URL"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Sender_Mail = os.Getenv("sender_mail")
	Sender_Password = os.Getenv("sender_password")
	DB_Password = os.Getenv("db_password")
	DB_Name = os.Getenv("db_name")
	DB_User = os.Getenv("db_user")
}
