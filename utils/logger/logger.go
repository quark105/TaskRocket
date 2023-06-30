package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.Formatter = &logrus.JSONFormatter{}
	Log.Level = logrus.DebugLevel
	Log.Out = os.Stdout
}
