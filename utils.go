package logger

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// GetLogWriter returns an io.Writer for log
// if already  created returns the same otherwise creates a new buffered Writer and returns
func GetLogWriter() (io.Writer, error) {
	file, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logDirectory, processName), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("[LOGGER]Error opening file", err.Error())
		return nil, err
	}
	return file, err
}

var _isTermial *bool

// IsTerminal checks if the stdout is a terminal or pipe
func IsTerminal() bool {
	if _isTermial != nil {
		return *_isTermial
	}
	_isTermial = isTerminal()
	return *_isTermial
}

func isTerminal() *bool {
	fd := int(os.Stdout.Fd())
	isT := terminal.IsTerminal(int(fd))
	return &isT
}
