package logger

import (
	"encoding/json"
	"os"
	"time"

	"github.com/spf13/viper"
)

//Log defines the structure of the log message or the log format
type Log struct {
	LogTime     time.Time `json:"time"`
	ProcessName string    `json:"process_name"`
	HostName    string    `json:"host_name"`
	ProcessID   int       `json:"process_id"`
	Level       string    `json:"level"`
	FileName    string    `json:"file_name"`
	LineNum     int       `json:"line_num"`
	Msg         string    `json:"log_msg"`
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
