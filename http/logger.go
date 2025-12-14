package http

import (
	"fmt"
	"log/slog"
)

type ClientLogger struct {
	l *slog.Logger
}

func NewClientLogger(l *slog.Logger) *ClientLogger {
	return &ClientLogger{l: l}
}

func (s *ClientLogger) Errorf(format string, args ...any) {
	s.l.Error("http", slog.String("msg", fmt.Sprintf(format, args...)))
}

func (s *ClientLogger) Warnf(format string, args ...any) {
	s.l.Warn("http", slog.String("msg", fmt.Sprintf(format, args...)))
}

func (s *ClientLogger) Debugf(format string, args ...any) {
	s.l.Debug("http", slog.String("msg", fmt.Sprintf(format, args...)))
}
