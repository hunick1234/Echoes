package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/hunick1234/Echoes/logger"
	sender "github.com/hunick1234/Echoes/sender/handler"
	"github.com/hunick1234/Echoes/server"
	httpb "github.com/hunick1234/Echoes/server/httpB"
	"github.com/hunick1234/Echoes/server/middleware"
)

func main() {
	router := httpb.WrappedMux{http.DefaultServeMux}
	server := server.Server{
		Addr: ":5050",
		Log: &logger.Log{
			Logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		},
		Handler: middleware.Logging(
			router.Next(
				middleware.End(),
			),
		),
	}
	sender.StartSenderHandle(&router)
	server.Start()
}
