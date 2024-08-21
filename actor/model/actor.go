package model

type Actor struct {
	Mail string `json:"mail"`
}

func (u *Actor) IsVaildMail() bool {
	return true
}
