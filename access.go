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
		Format: `{"time":"${time_rfc3339}","remote_addr":"${remote_ip}",` +
			`"request":"${method} ${uri} HTTP1.1",` +
			`"log_message" : [{"key" : "body_bytes_sent", "value" : "${bytes_out}"},` +
			`{"key" : "request_count", "value" : "true"},` +
			`{"key" : "status", "value" : "${status}"},` +
			`{"key" : "request_time", "value" : "${latency}"},` +
			`{"key" : "upstream_response_time", "value" : "${latency_human}"}],` +
			`"request_method":"${method}","uri":"${uri}",` +
			`"latency_human":"${latency_human}"}` + "\n",
		Output: file,
	}
	return &loggerConf, nil
}
