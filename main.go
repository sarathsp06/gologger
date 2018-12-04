package logger

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	logDirectory, processName, logLevel, hostName string
	processID                                     int
)

//Monitor accepts a metric and sends it across
func Monitor(m *Metric) {
	fmt.Fprint(os.Stdout, m)
}

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

//InitLogger initialise logger object with logWriter and log level
func InitLogger(level, directory, process string, humanRedable bool) error {
	if _logger != nil {
		return errors.New("Logger initiated already")
	}
	logDirectory, processName = directory, process
	logLevel, ok := LogLevels[level]
	if !ok {
		return errors.New("Invalid log level")
	}
	_log := new(Logger)
	_log.humanRedable = humanRedable
	logWriter, err := GetLogWriter()
	if err != nil {
		log.Println("Failed getting log writer :: ", err.Error())
		return err
	}
	_log.SetLogWriter(logWriter)
	_log.LogLevel = logLevel
	_logger = _log
	return nil
}

//SetBufferSize sets the buffer size for buffering logs
//the logs will be flushed anyways upon stop of the service
func SetBufferSize(bufferSize int) error {
	if _logger == nil {
		return errors.New("Nil logger ")
	}
	_logger.SetBufferSize(bufferSize)
	return nil
}

//GetLogger  returns the current default logger instance
func GetLogger() (*Logger, error) {
	if _logger == nil {
		return nil, errors.New("Nil logger ")
	}
	return _logger, nil
}

//SetLogWriter sets writer for log
func SetLogWriter(writer io.Writer) error {
	if writer == nil {
		return errors.New("Nil writer")
	}
	return _logger.SetLogWriter(writer)
}

//Flush flushes the data logs to log writer
func Flush() {
	_logger.Flush()
}

//SetLogType sets the log type
//This is more like a lable saying the type of log
//some possible log types would be application,error,
func SetLogType(logType string) {
	_logger.SetLogType(logType)
}
