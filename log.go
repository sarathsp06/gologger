package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

//Log defines the structure of the log message or the log format
type Log struct {
	ProcessName string `json:"service"`
	LogType     string `json:"logtype"`
	LogTime     string `json:"timestamp"`
	HostName    string `json:"host"`
	Msg         string `json:"message"`
	Level       string `json:"loglevel"`
	ProcessID   int    `json:"processid"`
	FileName    string `json:"filename"`
	LineNum     int    `json:"linenum"`
}

var logStruct *Log

//GetLog returns log struct
//GetLog returns the log struct with essential common data filled in
func GetLog() Log {
	if logStruct == nil {
		logStruct = new(Log)
		logStruct.ProcessName = processName
		logStruct.ProcessID = processID
		logStruct.HostName = hostName
	}
	logStruct.LogTime = time.Now().UTC().Format(time.RFC3339)
	return *logStruct
}

//String implements Stringer interface
//json encode the object and returns
func (log Log) String() string {
	marshalledData, err := json.Marshal(&log)
	if err != nil {
		return log.Msg
	}
	return string(marshalledData)
}

func logColor(logLevel string) int {
	if level, ok := LogColors[logLevel]; ok {
		return level
	}
	return WHITE
}

func fileNameColor() int {
	return WHITE
}

//Human returns human readable log  string
///Here the filename is going to be just to
func (log Log) Human() string {
	fileArray := strings.Split(log.FileName, string(os.PathSeparator))
	fileName := strings.Join(fileArray[len(fileArray)-2:], string(os.PathSeparator))
	humanFormat := "[%s]\033[%dm[%s]\t\033[%dm%s:%d~%s:%d\t\t\033[%dm%s\033[0m"
	logString := fmt.Sprintf(humanFormat,
		log.LogTime,
		logColor(log.Level),
		log.Level,
		fileNameColor(),
		hostName,
		processID,
		fileName,
		log.LineNum,
		logColor(log.Level),
		log.Msg)
	return logString
}

func init() {
	processID = os.Getpid()
	hostName, _ = os.Hostname()
}
