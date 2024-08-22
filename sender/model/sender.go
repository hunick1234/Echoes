package model

import (
	"bytes"
	"html/template"
	"net/smtp"
)

type Sender struct {
	mailAccount string
	password    string
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

func (e *EmailData) parseRegisterMailTemplate(templateFileName string) (string, error) {
	templateFileName = "mail_template.html"
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

func (s *Sender) FormatContent(body string) string {

	return ""
}

func (s *Sender) SendMail(to []string, subject, body string) error {
	from := s.mailAccount
	password := s.password

	// 設置SMTP服務器信息
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	// 設置電子郵件頭部和正文
	msg := "From: " + from + "\n" +
		"To: " + to[0] + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" + body

	// 設置認證
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// 發送郵件
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}
