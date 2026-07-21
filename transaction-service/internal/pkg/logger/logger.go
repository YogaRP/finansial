package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init(env string) {
	zerolog.TimeFieldFormat = time.RFC3339

	if env == "production" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}).With().Timestamp().Logger()
	}
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Infof(msg string, args ...any) {
	log.Info().Msgf(msg, args...)
}

func Debug(msg string) {
	log.Debug().Msg(msg)
}

func Debugf(msg string, args ...any) {
	log.Debug().Msgf(msg, args...)
}

func Warn(msg string) {
	log.Warn().Msg(msg)
}

func Warnf(msg string, args ...any) {
	log.Warn().Msgf(msg, args...)
}

func Error(msg string, err error) {
	log.Error().Err(err).Msg(msg)
}

func Errorf(msg string, err error, args ...any) {
	log.Error().Err(err).Msgf(msg, args...)
}

func Fatal(msg string, err error) {
	log.Fatal().Err(err).Msg(msg)
}

// WithField returns a logger with a custom field attached
// useful for attaching request ID, user ID, etc.
func WithField(key string, value any) zerolog.Logger {
	return log.With().Interface(key, value).Logger()
}
