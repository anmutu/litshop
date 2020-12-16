package logger

import (
	"fmt"
	"github.com/phachon/go-logger"
	"litshop/src/config"
	"litshop/src/pkg/path"
	"time"
)

var (
	logger    *go_logger.Logger
	asyncMode bool
)

func init() {
	logger = go_logger.NewLogger()
	asyncMode = false

	if config.GetBool("log.async") == true {
		asyncMode = true
	}

	logPath := path.StoragePath() + "/logs/"
	logFile := logPath + "app.log"

	fileConfig := &go_logger.FileConfig{
		Filename: logFile,
		LevelFileName: map[int]string{
			logger.LoggerLevel("error"): logPath + "error.log",
		},
		MaxSize:    1024 * 1024,
		MaxLine:    10000,
		DateSlice:  "d",
		JsonFormat: config.GetString("log.format") == "json",
		Format:     "%millisecond_format% [%level_string%] [%file%:%line%] %body%",
	}

	var level int = go_logger.LOGGER_LEVEL_DEBUG
	switch config.GetString("log.level") {
	case "info":
		level = go_logger.LOGGER_LEVEL_DEBUG
	case "warning":
		level = go_logger.LOGGER_LEVEL_WARNING
	case "error":
		level = go_logger.LOGGER_LEVEL_ERROR
	case "critical":
		level = go_logger.LOGGER_LEVEL_CRITICAL
	}

	_ = logger.Attach("file", level, fileConfig)
	if asyncMode {
		logger.SetAsync()

		go func() {
			for range time.Tick(time.Minute) {
				Flush()
			}
		}()
	}
}

func Emergency(msg string) {
	_ = logger.Writer(go_logger.LOGGER_LEVEL_EMERGENCY, msg)
}

func Emergencyf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	_ = logger.Writer(go_logger.LOGGER_LEVEL_EMERGENCY, msg)
}

func Alert(msg string) {
	_ = logger.Writer(go_logger.LOGGER_LEVEL_ALERT, msg)
}

//log alert format
func Alertf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	_ = logger.Writer(go_logger.LOGGER_LEVEL_ALERT, msg)
}

func Critical(msg string) {
	_ = logger.Writer(go_logger.LOGGER_LEVEL_CRITICAL, msg)
}

func Criticalf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	_ = logger.Writer(go_logger.LOGGER_LEVEL_CRITICAL, msg)
}

func Error(msg string) {
	_ = logger.Writer(go_logger.LOGGER_LEVEL_ERROR, msg)
}

func Warning(msg string) {
	_ = logger.Writer(go_logger.LOGGER_LEVEL_WARNING, msg)
}

func Warningf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	_ = logger.Writer(go_logger.LOGGER_LEVEL_WARNING, msg)
}

func Notice(msg string) {
	_ = logger.Writer(go_logger.LOGGER_LEVEL_NOTICE, msg)
}

func Noticef(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	_ = logger.Writer(go_logger.LOGGER_LEVEL_NOTICE, msg)
}

func Info(msg string) {
	_ = logger.Writer(go_logger.LOGGER_LEVEL_INFO, msg)
}

func Infof(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	_ = logger.Writer(go_logger.LOGGER_LEVEL_INFO, msg)
}

func Debug(msg string) {
	_ = logger.Writer(go_logger.LOGGER_LEVEL_DEBUG, msg)
}

func Debugf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	_ = logger.Writer(go_logger.LOGGER_LEVEL_DEBUG, msg)
}

func Flush() {
	if logger != nil && asyncMode {
		logger.Flush()
	}
}
