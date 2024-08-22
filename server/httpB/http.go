package httpb

import (
	"log"
	"net/http"
)

type WrappedMux struct {
	*http.ServeMux
}

func (wm *WrappedMux) Post(path string, handleFunc http.HandlerFunc) {
	restPath := "POST " + path
	wm.HandleFunc(restPath, handleFunc)
	log.Println("Register", restPath)
}

func (wm *WrappedMux) Get(path string, handleFunc http.HandlerFunc) {
	restPath := "GET " + path
	wm.HandleFunc(restPath, handleFunc)
	log.Println("Register", restPath)
}

func (wm *WrappedMux) Next(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wm.ServeHTTP(w, r)
		next.ServeHTTP(w, r)
	})
}
