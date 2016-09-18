package logger

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/viper"
)

var logWriter io.Writer

//GetLogWriter returns an io.Writer for log
//if already  created returns the same otherwise creates a new buffered Writer and returns
func GetLogWriter() (io.Writer, error) {
	if logWriter != nil {
		return logWriter, nil
	}
	file, err := os.OpenFile(fmt.Sprintf("%s/%s.log", viper.GetString("log_directory"), viper.GetString("process_name")), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Error opening file", err.Error())
		return nil, err
	}
	logWriter = bufio.NewWriterSize(file, 64)
	//return file, err
	return logWriter, err //buffer log with 1kB buffer
}
