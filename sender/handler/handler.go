package handler

import (
	"github.com/hunick1234/Echoes/sender/model"
	"github.com/hunick1234/Echoes/sender/repository"
	"github.com/hunick1234/Echoes/sender/service"
	httpb "github.com/hunick1234/Echoes/server/httpB"
	"github.com/hunick1234/Echoes/workpool"
)

func StartSenderHandle(router *httpb.WrappedMux) {
	// repo := repository.DefaultSenderRepo()
	repo := &repository.DefaultMock

	// TODO: Read the pool size from a configuration file or environment variable
	poolSize := 10
	wp := workpool.NewWorkerPool(poolSize)

	srv := service.NewSenderService(wp, repo)
	h := httpb.NewServiceHandler(srv, nil, nil)

	router.Post("/send-register-mail", httpb.Handler(SendRegisterMail, h))
}

func SendRegisterMail(s *httpb.ServiceHandler[service.SenderService]) error {
	mail := s.GetReq().URL.Query().Get("mail")
	err := s.GetSrv().SendMail(model.Register, []string{mail})
	if err != nil {
		return err
	}
	return nil
}
