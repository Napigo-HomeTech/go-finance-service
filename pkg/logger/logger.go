package logger

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func GetLogger() *log.Entry {
	return log.WithFields(log.Fields{
		"service":   os.Getenv("SERVICE_NAME"),
		"timestamp": time.Now(),
	})
}
