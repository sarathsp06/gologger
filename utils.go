package logger

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var logWriter io.Writer

//GetLogWriter returns an io.Writer for log
//if already  created returns the same otherwise creates a new buffered Writer and returns
func GetLogWriter() (io.Writer, error) {
	if logWriter != nil {
		return logWriter, nil
	}
	file, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logDirectory, processName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error opening file", err.Error())
		return nil, err
	}
	logWriter = bufio.NewWriterSize(file, 64)
	//return file, err
	return logWriter, err
}
