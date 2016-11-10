package main

import (
	"os"
	logger "github.com/sarathsp06/gologger"
)

func main() {
	if err := logger.InitLogger("INFO", ".", "sample_logger"); err != nil {
		panic(err.Error())
	}
	logger.SetLogWriter(os.Stdout)
	logger.Error("error happened")
	logger.Debug("Debug message")
	logger.Info("error happened")
	logger.Warning("error happened")
	logger.Errorf("error happened %s", "Yo")
	logger.Debugf("debug message : %s", "YoYo")
}
