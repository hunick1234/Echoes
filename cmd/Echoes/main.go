package main

import (
	"log/slog"
	"os"

	"github.com/hunick1234/Echoes/actor/handler"
	"github.com/hunick1234/Echoes/logger"
	"github.com/hunick1234/Echoes/server"
)

func main() {
	server := server.Server{
		Addr: ":8080",
		Log: &logger.Log{
			Logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
		},
	}
	handler.StartUserHAndle(&server)
	server.Start()
}
