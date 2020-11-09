package helpers

import (
	"os"

	log "github.com/sirupsen/logrus"
)

//Logger defines log profile
type Logger struct {
}

//Init setup Logger profile
func (l Logger) Init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(l.getLevel())
}

func (l Logger) getLevel() log.Level {
	appLogLevel := os.Getenv("LOG_LEVEL")
	switch appLogLevel {
	case "DEBUG":
		return log.DebugLevel
	case "INFO":
		return log.InfoLevel
	case "PRODUCTION":
		return log.ErrorLevel
	default:
		return log.TraceLevel
	}
}
