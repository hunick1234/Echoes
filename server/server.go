package server

import (
	"net/http"

	"github.com/hunick1234/Echoes/logger"
)

type Server struct {
	Addr string
	mux  http.ServeMux
	Log  *logger.Log
}

func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	s.Log.Info("Connect", "Method:", req.Method, "Url:", req.URL.Path)
	s.mux.ServeHTTP(res, req)
}

func (s *Server) Start() {
	s.Log.Debug("Starting server at port", "port", s.Addr)
	if err := http.ListenAndServe(s.Addr, s); err != http.ErrServerClosed {
		s.Log.Warn("ListenAndServe():", "err", err)
	}

}

func (s *Server) Post(path string, handleFunc http.HandlerFunc) {
	restPath := "POST " + path
	s.mux.HandleFunc(restPath, handleFunc)
	s.Log.Debug("Registered", "POST", path)
}

func (s *Server) Get(path string, handleFunc http.HandlerFunc) {
	restPath := "GET " + path
	s.mux.HandleFunc(restPath, handleFunc)
	s.Log.Info("Registered", "GET", path)
}
