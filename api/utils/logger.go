package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"strings"
)

type LoggerUtil struct{}

type Formatter struct {
	logrus.TextFormatter
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	format := "%s - %s - %s - %s\n"
	time := entry.Time.Format("2006-01-02 15:04:05")
	level := strings.ToUpper(entry.Level.String())
	file := "nil - nil - nil"
	if entry.HasCaller() {
		file = fmt.Sprintf("%s() - %d", entry.Caller.Function, entry.Caller.Line)
	}
	return []byte(fmt.Sprintf(format, time, level, file, entry.Message)), nil
}

func (l LoggerUtil) Logger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(new(Formatter))
	log.SetReportCaller(true)
	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/password-go-api.log", // Log file path
		MaxSize:    100,                        // Maximum size saved for each log file Unit: M
		MaxBackups: 30,                         // The maximum number of backups that can be saved in a log file
		MaxAge:     180,                        // The maximum number of days a file can be saved
		Compress:   true,                       // Whether to compress
	})
	log.SetLevel(logrus.InfoLevel)
	return log
}
