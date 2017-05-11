package logger

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
)

//Logger struct to hold the log level and the Writer
type Logger struct {
	LogLevel        int
	BufferSizeBytes int
	LogType         string
	humanRedable    bool
	writer          io.Writer
	bufferedWriter  io.Writer
}

var _logger *Logger

//SetLogType sets the log type
//This is more like a lable saying the type of log
//some possible log types would be application,error,
func (l *Logger) SetLogType(logType string) *Logger {
	l.LogType = logType
	return l
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
	logStruct.LogType = l.LogType
	if l.humanRedable {
		l.Writer().Write([]byte(logStruct.Human() + "\n\r"))
		return
	}
	l.Writer().Write([]byte(logStruct.String() + "\n\r"))
	return
}

//Flush flushed the buffer
func (l Logger) Flush() {
	if writer, ok := l.bufferedWriter.(*bufio.Writer); ok {
		fmt.Println("Flushing log buffer")
		if err := writer.Flush(); err != nil {
			fmt.Fprintln(os.Stderr, "Failed flushing the logs", err.Error())
		}
	}
}

//SetBufferSize sets buffer size as bytes
func (l *Logger) SetBufferSize(bufferSizeBytes int) (err error) {
	l.BufferSizeBytes = bufferSizeBytes
	l.bufferedWriter = l.writer
	if bufferSizeBytes > 0 {
		l.bufferedWriter = bufio.NewWriterSize(l.writer, bufferSizeBytes)
	}
	return nil
}

//Writer returns writer to be used
func (l *Logger) Writer() io.Writer {
	return l.bufferedWriter
}

//SetLogWriter sets default  writer
func (l *Logger) SetLogWriter(writer io.Writer) error {
	if writer == nil {
		return errors.New("Nil writer")
	}
	l.Flush()
	l.writer = writer
	l.SetBufferSize(l.BufferSizeBytes)
	return nil
}
