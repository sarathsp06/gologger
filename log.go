package logger

import (
	"encoding/json"
	"os"
	"time"

	"github.com/spf13/viper"
)

//Log defines the structure of the log message or the log format
type Log struct {
	LogTime     time.Time
	ProcessName string
	HostName    string
	ProcessID   int
	Level       string
	FileName    string
	LineNum     int
	Msg         string
}

var logStruct *Log

//GetLog returns log struct
//GetLog returns the log struct with essential common data filled in
func GetLog() Log {
	if logStruct == nil {
		logStruct = new(Log)
		logStruct.ProcessName = viper.GetString("process_name")
		logStruct.ProcessID = os.Getpid()
		logStruct.HostName, _ = os.Hostname()
	}
	logStruct.LogTime = time.Now()
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
