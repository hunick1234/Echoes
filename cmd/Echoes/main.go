package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/hunick1234/Echoes/actor/handler"
	"github.com/hunick1234/Echoes/logger"
	"github.com/hunick1234/Echoes/server"
	httpb "github.com/hunick1234/Echoes/server/httpB"
	"github.com/hunick1234/Echoes/server/middleware"
)

func main() {
	router := httpb.WrappedMux{http.NewServeMux()}
	server := server.Server{
		Addr: ":8080",
		Log: &logger.Log{
			Logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		},
		Handler: middleware.Logging(
			router.Next(
				middleware.Cors(
					middleware.End(),
				),
			),
		),
	}

	handler.StartUserHAndle(&router)
	server.Start()
}
