package model

import (
	"bytes"
	"html/template"
	"os"

	"github.com/hunick1234/Echoes/config"
)

var (
	from = config.Sender_Mail
)

type WorkType int

const (
	Register WorkType = iota
	Resend
)

type Sender struct {
	MailAccount string
	Passworld   string
	EmailFormate
	EmailData
}

type EmailFormate struct {
	Subject string
	Body    string
	To      []string
}

type EmailData struct {
	Name    string
	Message string
	Url     string
}

func (e *EmailData) ParseRegisterMailTemplate() (string, error) {
	templateFileName := "../../template/register_mail_template.html"
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, e); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func DefaultSender() Sender {
	return Sender{
		MailAccount: os.Getenv("sender_mail"),
		Passworld:   os.Getenv("sender_passworld"),
		EmailFormate: EmailFormate{
			Subject: "test",
			To:      []string{"hunick1234@gmail.com"},
			Body:    "",
		},
		EmailData: EmailData{},
	}
}

func SetFormateMail(subject, body string) string {
	msg := "From: " + from + "\n" +
		"To: no-reply@example.com\n" +
		"Subject: " + subject + "\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" + body
	return msg
}
