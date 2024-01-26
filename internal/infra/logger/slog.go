package logger

import (
	"log/slog"
	"os"

	"gopaseto/internal/core/config"

	slogmulti "github.com/samber/slog-multi"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *slog.Logger

func Set(conf *config.APP) {
	logger = slog.New(
		slog.NewTextHandler(os.Stderr, nil),
	)

	if conf.Env == "production" {
		logRotate := &lumberjack.Logger{
			Filename:   "log/app.log",
			MaxSize:    100, //mb
			MaxBackups: 3,
			MaxAge:     28, //days
			Compress:   true,
		}

		logger = slog.New(
			slogmulti.Fanout(
				slog.NewJSONHandler(logRotate, nil),
				slog.NewTextHandler(os.Stderr, nil),
			),
		)
	}

	slog.SetDefault(logger)
}
