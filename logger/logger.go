package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger *zerolog.Logger
}

func NewLogger(level string) *Logger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	skipFrameCount := 3
	logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &Logger{
		logger: &logger,
	}
}

func (l *Logger) Debug(msg string) {
	l.logger.Debug().Msg(msg)

}

func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)

}

func (l *Logger) Warn(msg string) {
	l.logger.Warn().Msg(msg)

}

func (l *Logger) Error(err error, msg string) {
	l.logger.Err(err).Msg(msg)
}
