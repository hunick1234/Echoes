package handler

import (
	"net/http"

	"github.com/hunick1234/Echoes/sender/repository"
	"github.com/hunick1234/Echoes/sender/service"
	"github.com/hunick1234/Echoes/server"
)

type SenderHandler struct {
	srv service.SenderService
	res http.ResponseWriter
	req *http.Request
}

func NewSenderHandler(srv service.SenderService, res http.ResponseWriter, req *http.Request) SenderHandler {
	return SenderHandler{
		srv: srv,
		res: res,
		req: req,
	}
}

func StartSenderHandle(ser *server.Server) {
	ser.Get("/SendRegisterMail", handler(SendRegisterMail))
	ser.Get("/SendMail",handler(SendMail))
}

func handler(api func(*SenderHandler)) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		repo := repository.NewSenderRepo()
		srv := service.NewSenderService(repo)
		s := NewSenderHandler(srv, res, req)
		api(&s)
	}
}

func SendRegisterMail(s *SenderHandler) {

}

func SendMail(s *SenderHandler) {

}
