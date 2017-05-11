package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

//GetLogWriter returns an io.Writer for log
//if already  created returns the same otherwise creates a new buffered Writer and returns
func GetLogWriter() (io.Writer, error) {
	file, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logDirectory, processName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("[LOGGER]Error opening file", err.Error())
		return nil, err
	}
	return file, err
}
