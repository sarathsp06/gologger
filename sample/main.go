package main

import (
	"bitbucket.org/enterpriseproducts/common/go/logger"
	"os"
)

func main() {
	if err := logger.InitLogger("DEBUG", ".", "sample_logger", true); err != nil {
		panic(err.Error())
	}
	logger.SetLogWriter(os.Stdout)
	logger.Monitor(logger.NewMetric("test"))
	logger.SetLogType("application")
	logger.Error("error happened")
	logger.Debug("Debug message")
	logger.Info("error happened")
	logger.Warning("error happened")
	logger.Errorf("error happened %s", "Yo")
	logger.Debugf("debug message : %s", "YoYo")
	logger.Flush()
}
