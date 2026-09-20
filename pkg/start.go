package pkg

import (
	"log/slog"
	"os"
)

func Start(fn func(testing bool) string) {
	testing := false

	for _, arg := range os.Args[1:] {
		switch arg {
		case "--debug":
			slog.SetLogLoggerLevel(slog.LevelDebug)
		case "--testing":
			testing = true
		default:
			slog.Warn("unknown argument", "arg", arg)
			os.Exit(1)
		}
	}

	slog.Info("finished", "result", fn(testing))
}
