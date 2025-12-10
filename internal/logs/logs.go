package logs

import (
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLog(level logrus.Level, formatter logrus.Formatter) {
	Log = logrus.New()
	Log.SetLevel(level)
	Log.SetFormatter(formatter)
}
