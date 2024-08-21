package main

import (
	"log/slog"

	"github.com/hunick1234/Echoes/logger"
	sender "github.com/hunick1234/Echoes/sender/handler"
	"github.com/hunick1234/Echoes/server"
)

func main() {
	server := server.Server{
		Addr: ":5050",
		Log: &logger.Log{
			Logger: slog.Default(),
		},
	}
	sender.StartSenderHandle(&server)
	server.Start()
}
