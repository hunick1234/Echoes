package logger

import (
	"log/slog"
	"os"
)

var DefaultLog = slog.New(slog.NewJSONHandler(os.Stdout, nil))

type Log struct {
	*slog.Logger
}
