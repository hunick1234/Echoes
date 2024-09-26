package model

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"regexp"
	"time"
)

type Actor struct {
	Mail string `json:"mail", `
}

type RegisterActor struct {
	tableName  struct{}  `pg:"user"`
	Id         int32     `pg:"id"`
	Mail       string    `pg:"mail,unique	"`
	MailId     string    `pg:"mail_id,unique"`
	CreateTime time.Time `pg:"default:now()"`
	UpdateTime time.Time `pg:"update_time"`
}

func GenerateID(email string) string {
	t := time.Now().Unix() // 使用時間戳來定期變動
	data := fmt.Sprintf("%s-%d", email, t)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

func ValidateEmail(email string) error {
	if !regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`).MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}
