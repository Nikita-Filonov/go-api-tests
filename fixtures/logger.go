package fixtures

import (
	"log/slog"
	"os"

	"github.com/Nikita-Filonov/axiom"
)

func SetLoggerFixture(_ *axiom.Config) (any, func(), error) {
	logger := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
	)

	return logger, nil, nil
}

func GetLoggerFixture(cfg *axiom.Config) *slog.Logger {
	return axiom.GetFixture[*slog.Logger](cfg, "logger")
}
