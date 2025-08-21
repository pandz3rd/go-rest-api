package app

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	level := fmt.Sprintf("%-5s", entry.Level.String())

	traceId, ok := entry.Data["traceId"]
	if !ok {
		traceId = "-"
	}

	filename := "unknown"
	if entry.Caller != nil {
		filename = filepath.Base(entry.Caller.File)
	}

	logLine := fmt.Sprintf("[%s] [%s] [%s:%d] [%s] - %s",
		timestamp,
		level,
		filename,
		entry.Caller.Line,
		traceId,
		entry.Message,
	)

	// Print other fields (skip class & traceId)
	for k, v := range entry.Data {
		if k == "class" || k == "traceId" {
			continue
		}
		logLine += fmt.Sprintf(" %s=%v", k, v)
	}

	logLine += "\n"
	return []byte(logLine), nil
}

func InitLogger() *logrus.Logger {
	logger := logrus.New()

	rotateLogger := &lumberjack.Logger{
		Filename:   "./log/app.log", // always writes to this, old files get renamed with timestamps
		MaxSize:    10,              // MB before rotation
		MaxBackups: 24,              // keep up to 24 backups
		MaxAge:     7,               // keep 7 days
		Compress:   true,            // compress old logs
	}

	logger.SetOutput(io.MultiWriter(os.Stdout, rotateLogger))
	logger.SetFormatter(&CustomFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	logger.SetReportCaller(true)

	return logger
}
