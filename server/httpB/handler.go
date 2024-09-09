package httpb

import (
	"net/http"
)

type ServiceHandler[T any] struct {
	srv T
	req *http.Request
	res http.ResponseWriter
}

func NewServiceHandler[T any](srv T, req *http.Request, res http.ResponseWriter) *ServiceHandler[T] {
	return &ServiceHandler[T]{
		srv: srv,
		req: req,
		res: res,
	}
}

func (s *ServiceHandler[T]) GetSrv() T {
	return s.srv
}

func (s *ServiceHandler[T]) GetReq() *http.Request {
	return s.req
}

func (s *ServiceHandler[T]) GetRes() http.ResponseWriter {
	return s.res
}

func (s *ServiceHandler[T]) SetRes(res http.ResponseWriter) *ServiceHandler[T] {
	s.res = res
	return s
}

func (s *ServiceHandler[T]) SetReq(req *http.Request) *ServiceHandler[T] {
	s.req = req
	return s
}

func (s *ServiceHandler[T]) SetSrv(srv T) *ServiceHandler[T] {
	s.srv = srv
	return s
}

func Handler[T any](api func(*ServiceHandler[T]) error, h *ServiceHandler[T]) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		h.SetRes(res).SetReq(req)
		err := api(h)
		if err != nil {

		}
	}
}
