package logger

import (
	"github.com/TeplyyMaksim/bookstore_users-api/utils/errors_utils"
	"github.com/rs/zerolog"
	"os"
)

var (
	log zerolog.Logger
)

func init () {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log = zerolog.New(os.Stderr).With().Timestamp().Logger()
}

func Info (msg string) {
	log.Info().Msg(msg)
}

func Error (msg string, ) {
	log.Error().Msg(msg)
}

func HttpError (httpError *errors_utils.HttpError) {
	log.Error().
		Int("status", httpError.Status).
		Str("error", httpError.Error).
		Msg(httpError.Message)
}

func GetLogger() zerolog.Logger {
	return log
}