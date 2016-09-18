package logger

import "github.com/labstack/echo/middleware"

//GetEchoLoggerConfiguration creates a LoggerConfig for echo middleware and returns it
//Uses the same writer as other logs uses
func GetEchoLoggerConfiguration() (*middleware.LoggerConfig, error) {
	file, err := GetLogWriter()
	if err != nil {
		return nil, err
	}
	loggerConf := middleware.LoggerConfig{
		Format: `{"LogTime":"${time_rfc3339}","RemoteIP":"${remote_ip}",` +
			`"Method":"${method}","URI":"${uri}","Status":${status}, "Latency":${latency},` +
			`"LatencyHuman":"${latency_human}","RxBytes":${rx_bytes},` +
			`"TxBytes":${tx_bytes}}` + "\n",
		Output: file,
	}
	return &loggerConf, nil
}
