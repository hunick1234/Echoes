package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/hunick1234/Echoes/actor/model"
	"github.com/hunick1234/Echoes/actor/repository"
	"github.com/hunick1234/Echoes/actor/service"
	httpb "github.com/hunick1234/Echoes/server/httpB"
)

func StartUserHandle(router *httpb.WrappedMux) {
	repo := repository.DefaultActorRepo()
	srv := service.NewActorService(repo)
	h := httpb.NewServiceHandler(srv, nil, nil)

	router.Handle("/", RegisterPage("/"))
	router.Get("/register", httpb.Handler(Register, h))
	router.Get("/resend-mail", httpb.Handler(Resend, h))
}

func RegisterPage(prefix string) http.Handler {
	staticDir := http.Dir("../../web/actor")
	fileServer := http.FileServer(staticDir)
	return http.StripPrefix(prefix, fileServer)
}

func Register(u *httpb.ServiceHandler[service.ActorService]) error {
	mail := u.GetReq().URL.Query().Get("mail")

	//check then register
	if ok := u.GetSrv().CheckByMail(mail); !ok {
		registerActor := &model.RegisterActor{
			Mail:   mail,
			MailId: model.GenerateID(mail),
		}
		err := u.GetSrv().CreatUser(registerActor)
		if err != nil {
			return err
		}
	}

	ctx, cancel := context.WithTimeout(u.GetReq().Context(), 2*time.Second)
	defer cancel()

	sendMailUrl := "http://127.0.0.1:5050/send-register-mail?" + mail
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, sendMailUrl, nil)
	if err != nil {
		return err
	}

	// 發送 HTTP 請求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func Resend(a *httpb.ServiceHandler[service.ActorService]) error {
	return nil
}
