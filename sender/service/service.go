package service

import (
	"net/smtp"
	"os"

	"github.com/hunick1234/Echoes/sender/repository"
)

type SenderService interface {
		SendMail(to []string, subject, body string) error
}

func NewSenderService(repo *repository.SendRepo) SenderService {
	return &SenderServiceImpl{
		repo: repo,
	}
}

type SenderServiceImpl struct {
	repo *repository.SendRepo
}

func (s *SenderServiceImpl) SendMail(to []string, subject, body string) error {
	passworld := os.Getenv("sender_passworld")
	from := os.Getenv("sender_mail")

	// 設置SMTP服務器信息
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// 設置電子郵件頭部和正文
	msg := "From: " + from + "\n" +
		"To: " + to[0] + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" + body

	// 設置認證
	auth := smtp.PlainAuth("", from, passworld, smtpHost)

	// 發送郵件
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}
