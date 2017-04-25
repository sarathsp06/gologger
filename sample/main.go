package main

import (
	logger "github.com/sarathsp06/gologger"
	"os"
)

func main() {
	if err := logger.InitLogger("DEBUG", ".", "sample_logger", false); err != nil {
		panic(err.Error())
	}
        logger.SetLogType("application")
	logger.SetLogWriter(os.Stdout)
	logger.Error("error happened")
	logger.Debug("Debug message")
	logger.Info("error happened")
	logger.Warning("error happened")
	logger.Errorf("error happened %s", "Yo")
	logger.Debugf("debug message : %s", "YoYo")
}
