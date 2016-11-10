package logger


//ILogger interface defining the log functions
type ILogger interface {
	Log(depth int, level string, message string)
	Flush()
}

var _logger ILogger

