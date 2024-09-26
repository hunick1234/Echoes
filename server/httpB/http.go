package httpb

import (
	"encoding/json"
	"net/http"

	"github.com/hunick1234/Echoes/logger"
)

type WrappedMux struct {
	*http.ServeMux
}

type WrappedRepWriter struct {
	StatusCode int
	http.ResponseWriter
}

func (wm *WrappedMux) Post(path string, handleFunc http.HandlerFunc) {
	restPath := "POST " + path
	wm.HandleFunc(restPath, handleFunc)
	logger.DefaultLog.Info("Register", "restPath", restPath)
}

func (wm *WrappedMux) Get(path string, handleFunc http.HandlerFunc) {
	restPath := "GET " + path
	wm.HandleFunc(restPath, handleFunc)
	logger.DefaultLog.Info("Register", "restPath", restPath)
}

func (wm *WrappedMux) Next(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wm.ServeHTTP(w, r)
		next.ServeHTTP(w, r)
	})
}

func HttpResponJson(w http.ResponseWriter, statusCode int, message []byte) {
	// write respon to json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(message)
	w.Write(message)

}

func HttpRespon(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}
