package logging

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const (
	LevelTrace   = 0
	LevelDebug   = 1
	LevelConfig  = 2
	LevelInfo    = 3
	LevelWarning = 4
	LevelError   = 5
	LevelFatal   = 6
)

func levelToString(level int) string {
	if level > 6 || level < 0 {
		return "UNKNWON"
	} else {
		return levelNames[level]
	}
}

var logger *log.Logger
var logLevel = 3
var levelNames map[int]string
var format string

func init() {
	logger = log.New(os.Stdout, "", 0)
	levelNames = make(map[int]string)
	levelNames[0] = "TRACE"
	levelNames[1] = "DEBUG"
	levelNames[2] = "CONFIG"
	levelNames[3] = "INFO"
	levelNames[4] = "WARNING"
	levelNames[5] = "ERROR"
	levelNames[6] = "FATAL"

	format = "[{LEVEL}] {FILE}:{LINE}: {MSG}"
}

func SetLevel(level int) {
	logLevel = level
}

func formatString(level int, msg, file string, line int) string {
	s := strings.Replace(format, "{LEVEL}", levelToString(level), -1)

	s = strings.Replace(s, "{MSG}", msg, -1)

	s = strings.Replace(s, "{FILE}", file, -1)

	s = strings.Replace(s, "{LINE}", strconv.FormatInt(int64(line), 10), -1)

	return s
}

func log_(level int, format string, v ...interface{}) {
	if level < logLevel {
		return
	}

	msg := fmt.Sprintf(format, v...)
	_, file, line, _ := runtime.Caller(2)
	wd, err := os.Getwd()
	if err == nil && len(wd) < len(file) {
		file = file[len(wd):]
	}

	logger.Println(formatString(level, msg, file, line))
}

type Logger struct {
}

func (l Logger) Log(level int, format string, v ...interface{}) {
	log_(level, format, v...)
}

func (l Logger) Trace(format string, v ...interface{}) {
	log_(LevelTrace, format, v...)
}

func (l Logger) Debug(format string, v ...interface{}) {
	log_(LevelDebug, format, v...)
}

func (l Logger) Info(format string, v ...interface{}) {
	log_(LevelInfo, format, v...)
}

func (l Logger) Warning(format string, v ...interface{}) {
	log_(LevelWarning, format, v...)
}

func (l Logger) Errorf(format string, v ...interface{}) {
	log_(LevelError, format, v...)
}
