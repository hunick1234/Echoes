package handler

import (
	"net/http"

	"github.com/hunick1234/Echoes/actor/repository"
	"github.com/hunick1234/Echoes/actor/service"
	httpb "github.com/hunick1234/Echoes/server/httpB"
)

type ActorHandler struct {
	Service service.ActorService
	res     http.ResponseWriter
	req     *http.Request
}

func NewUserHandler(srv service.ActorService, res http.ResponseWriter, req *http.Request) ActorHandler {
	return ActorHandler{
		Service: srv,
		res:     res,
		req:     req,
	}
}

func StartUserHAndle(router *httpb.WrappedMux) {
	router.Get("/register", handler(Register))
}

func handler(api func(*ActorHandler)) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		repo := repository.NewActorRepo()
		srv := service.NewActorService(&repo)
		u := NewUserHandler(srv, res, req)
		api(&u)

	}
}

func Register(u *ActorHandler) {
	u.Register()
	http.Get("http://127.0.0.1:5050/SendRegisterMail")
}
