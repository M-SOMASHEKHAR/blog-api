package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func LoadLogger() {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	Log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Caller().Timestamp().Logger()
}
