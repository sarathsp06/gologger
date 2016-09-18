package logger

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

const (
	//MONITOR monitoring logs flag
	MONITOR string = "MONITOR"
	//ERROR error logs flag
	ERROR string = "ERROR"
	//WARNING Warning logs flag string
	WARNING string = "WARNING"
	//INFO informative log message flag string
	INFO string = "INFO"
	//DEBUG debug logs flag string
	DEBUG string = "DEBUG"
)

//LogLevels defines loglevel priorities  0 highest and 3 lowest
var LogLevels = map[string]int{MONITOR: 0, ERROR: 1, WARNING: 2, INFO: 3, DEBUG: 4}

//ILogger interface defining the log functions
type ILogger interface {
	Log(depth int, level string, message ...interface{})
}

//Logger struct to hold the log level and the Writer
type Logger struct {
	LogLevel int
	Writer   io.Writer
}

var _logger *Logger

//Debug Debug log without formatting
func Debug(message ...interface{}) {
	_logger.Log(1, DEBUG, fmt.Sprint(message...))
}

//Error function for error logs without formatting
func Error(message ...interface{}) {
	_logger.Log(1, ERROR, fmt.Sprint(message...))
}

//Info info level logs without formatting
func Info(message ...interface{}) {
	_logger.Log(1, INFO, fmt.Sprint(message...))
}

//Warning Warning level logs without formatting
func Warning(message ...interface{}) {
	_logger.Log(1, WARNING, fmt.Sprint(message...))
}

//Errorf Prints log with formatting
func Errorf(message ...interface{}) {
	_logger.Log(1, ERROR, fmt.Sprintf(message[0].(string), message[1:]...))
}

//Debugf Prints log with formatting
func Debugf(message ...interface{}) {
	_logger.Log(1, DEBUG, fmt.Sprintf(message[0].(string), message[1:]...))
}

//Warningf Prints log with formatting
func Warningf(message ...interface{}) {
	_logger.Log(1, WARNING, fmt.Sprintf(message[0].(string), message[1:]...))
}

//Infof Prints log with formatting
func Infof(message ...interface{}) {
	_logger.Log(1, INFO, fmt.Sprintf(message[0].(string), message[1:]...))
}

//Log given the stack depth and level with an array of messages
//decides if to be Written to logs ans writes to log with FileName and LineNum
//taken from runtime Info
func (l Logger) Log(depth int, level string, message string) {
	level = strings.ToUpper(level)
	if levelPriority, ok := LogLevels[level]; ok && levelPriority > l.LogLevel {
		return
	}

	_, file, line, _ := runtime.Caller(depth + 1)

	logStruct := GetLog()
	logStruct.Msg = message
	logStruct.FileName = file
	logStruct.LineNum = line
	logStruct.Level = level
	l.Writer.Write([]byte(logStruct.String() + "\n\r"))
}

//InitLogger initialise logger object with logWriter and log level
func InitLogger() error {
	_logger = &Logger{}
	logWriter, err := GetLogWriter()
	if err != nil {
		log.Println("Failed getting log writer", err.Error())
		return err
	}
	_logger.Writer = logWriter
	_logger.LogLevel = LogLevels[viper.GetString("log_level")]
	return nil
}

//Flush flushes the data logs to log writer
func Flush() {
	if writer, ok := _logger.Writer.(*bufio.Writer); ok {
		fmt.Println("Flushing log buffer")
		if err := writer.Flush(); err != nil {
			fmt.Fprintln(os.Stderr, "Failed flushing the logs", err.Error())
		}
	}
}
