package server

import (
	"net/http"

	"github.com/hunick1234/Echoes/logger"
)

type Server struct {
	Addr string
	Log  *logger.Log
	http.Handler
}

type WrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//s.Log.Info("Connect", "Method:", req.Method, "Url:", req.URL.Path)
	// s.mux.ServeHTTP(res, req)
}

func (s *Server) Start() {
	s.Log.Debug("Starting server at port", "port", s.Addr)
	if err := http.ListenAndServe(s.Addr, s.Handler); err != http.ErrServerClosed {
		s.Log.Warn("ListenAndServe():", "err", err)
	}
}
