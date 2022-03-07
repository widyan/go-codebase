package helper

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CustomLogger struct {
	logger *logrus.Logger
}

func CreateLogger(log *logrus.Logger) *CustomLogger {
	return &CustomLogger{log}
}

type Fields logrus.Fields

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	} else {
		slash := strings.LastIndex(file, "/")
		if slash >= 0 {
			file = file[slash+1:]
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

// Debug logs a message at level Debug on the standard logger.
func (log *CustomLogger) Debug(args ...interface{}) {
	if log.logger.Level >= logrus.DebugLevel {
		entry := log.logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Data["header"] = ""
		entry.Data["body"] = ""
		entry.Debug(args)
	}
}

// Debug logs a message with fields at level Debug on the standard logger.
func (log *CustomLogger) DebugWithFields(l interface{}, f Fields) {
	if log.logger.Level >= logrus.DebugLevel {
		entry := log.logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Data["header"] = ""
		entry.Data["body"] = ""
		entry.Debug(l)
	}
}

// Info logs a message at level Info on the standard logger.
func (log *CustomLogger) Info(args ...interface{}) {
	if log.logger.Level >= logrus.InfoLevel {
		entry := log.logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Data["header"] = ""
		entry.Data["body"] = ""
		entry.Info(args...)
	}
}

// Debug logs a message with fields at level Debug on the standard logger.
func (log *CustomLogger) InfoWithFields(l interface{}, f Fields) {
	if log.logger.Level >= logrus.InfoLevel {
		entry := log.logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Data["header"] = ""
		entry.Data["body"] = ""
		entry.Info(l)
	}
}

// Warn logs a message at level Warn on the standard logger.
func (log *CustomLogger) Warn(args ...interface{}) {
	if log.logger.Level >= logrus.WarnLevel {
		entry := log.logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Data["header"] = ""
		entry.Data["body"] = ""
		entry.Warn(args...)
	}
}

// Debug logs a message with fields at level Debug on the standard logger.
func (log *CustomLogger) WarnWithFields(l interface{}, f Fields) {
	if log.logger.Level >= logrus.WarnLevel {
		entry := log.logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Data["header"] = ""
		entry.Data["body"] = ""
		entry.Warn(l)
	}
}

// Error logs a message at level Error on the standard logger.
func (log *CustomLogger) Error(args ...interface{}) {
	if log.logger.Level >= logrus.ErrorLevel {
		entry := log.logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Data["header"] = ""
		entry.Data["body"] = ""
		entry.Error(args...)
	}
}

func (log *CustomLogger) ErrorWithContext(c *gin.Context, body interface{}, args ...interface{}) {
	if log.logger.Level >= logrus.ErrorLevel {
		b := make(map[string]interface{})
		for k, v := range c.Request.Header {
			b[k] = v
		}

		strHeader := ""
		header, err := json.Marshal(b)
		if err != nil {
			strHeader = ""
		} else {
			strHeader = string(header)
		}

		strBody := ""
		if body != nil {
			body, err := json.Marshal(body)
			if err != nil {
				strBody = ""
			} else {
				strBody = string(body)
			}
		}

		entry := log.logger.WithFields(logrus.Fields{})
		entry.Data["header"] = strHeader
		entry.Data["body"] = strBody
		entry.Data["file"] = fileInfo(2)
		entry.Error(args...)
	}
}

// Debug logs a message with fields at level Debug on the standard logger.
func (log *CustomLogger) ErrorWithFields(l interface{}, f Fields) {
	if log.logger.Level >= logrus.ErrorLevel {
		entry := log.logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Data["header"] = ""
		entry.Data["body"] = ""
		entry.Error(l)
	}
}

// Fatal logs a message at level Fatal on the standard logger.
func (log *CustomLogger) Fatal(args ...interface{}) {
	if log.logger.Level >= logrus.FatalLevel {
		entry := log.logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Data["header"] = ""
		entry.Data["body"] = ""
		entry.Fatal(args...)
	}
}

// Debug logs a message with fields at level Debug on the standard logger.
func (log *CustomLogger) FatalWithFields(l interface{}, f Fields) {
	if log.logger.Level >= logrus.FatalLevel {
		entry := log.logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Data["header"] = ""
		entry.Data["body"] = ""
		entry.Fatal(l)
	}
}

// Panic logs a message at level Panic on the standard logger.
func (log *CustomLogger) Panic(args ...interface{}) {
	if log.logger.Level >= logrus.PanicLevel {
		entry := log.logger.WithFields(logrus.Fields{})
		entry.Data["file"] = fileInfo(2)
		entry.Data["header"] = ""
		entry.Data["body"] = ""
		entry.Panic(args...)
	}
}

// Debug logs a message with fields at level Debug on the standard logger.
func (log *CustomLogger) PanicWithFields(l interface{}, f Fields) {
	if log.logger.Level >= logrus.PanicLevel {
		entry := log.logger.WithFields(logrus.Fields(f))
		entry.Data["file"] = fileInfo(2)
		entry.Data["header"] = ""
		entry.Data["body"] = ""
		entry.Panic(l)
	}
}
