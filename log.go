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
	LogTime     string `json:"time"`
	ProcessName string `json:"process_name"`
	HostName    string `json:"host_name"`
	ProcessID   int    `json:"process_id"`
	Level       string `json:"level"`
	FileName    string `json:"file_name"`
	LineNum     int    `json:"line_num"`
	Msg         string `json:"log_msg"`
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

//Human returns human readable log  string
///Here the filename is going to be just to
func (log Log) Human() string {
	fileArray := strings.Split(log.FileName, string(os.PathSeparator))
	fileName := strings.Join(fileArray[len(fileArray)-2:], string(os.PathSeparator))
	humanFormat := "[%s][%s]\t%s:%d %s:%d\t%s"
	logString := fmt.Sprintf(humanFormat,
		log.LogTime,
		log.Level,
		hostName,
		processID,
		fileName,
		log.LineNum,
		log.Msg)
	return logString
}

func init() {
	processID = os.Getpid()
	hostName, _ = os.Hostname()
}
