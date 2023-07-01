package logger

import (
	log "github.com/sirupsen/logrus"
)

var ServiceLogger *log.Entry

type LoggingConfig struct {
	ServiceName string `mapstructure:"service"`
	Formatter   string `mapstructure:"formatter"`
	Level       string `mapstructure:"level"`
}

func ConfigureLogger(lc LoggingConfig) {
	// Text formatter is set by default
	if lc.Formatter == "json" {
		log.SetFormatter(&log.JSONFormatter{})
	}

	// INFO level is set by default
	switch lc.Level {
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	}

	ServiceLogger = log.WithField("service", lc.ServiceName)
}

func init() {
	ServiceLogger = log.WithField("application", "TaskRocket")
}
