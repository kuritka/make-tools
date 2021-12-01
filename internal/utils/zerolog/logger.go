// Package log wraps zerolog logger and provides
// standard log functionality
package zerolog

import (
	"os"

	"github.com/rs/zerolog"
)

// Log is the global logger.
var Log *zerolog.Logger

//Logger returns the global logger.
func Logger() *zerolog.Logger {
	return Log
}

//init initializes the logger
func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: false} ).
		With().
		Timestamp().
		Logger()

	Log = &l
}
