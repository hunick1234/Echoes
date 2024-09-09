package service

import (
	"context"
	"net/smtp"

	"github.com/hunick1234/Echoes/config"
	"github.com/hunick1234/Echoes/workpool"
)

var (
	from     = config.Sender_Mail
	password = config.Sender_Password
)

//task
//register mail
//resend mail
//send echoes mail

func creatRegisterMailTask(to []string, msg string) workpool.Jober {
	t := func(ctx context.Context) error {
		err := sendMail(to, msg)
		if err != nil {
			return err
		}
		return nil
	}
	return workpool.CreatJob(t)
}

func sendMail(to []string, msg string) error {
	// 設置SMTP服務器信息
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// 設置認證
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// 發送郵件
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}
