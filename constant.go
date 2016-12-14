package logger

const (
	//MONITOR monitoring logs flag
	MONITOR string = "MONITOR"
	//ERROR error logs flag
	ERROR string = "ERROR"
	//WARNING Warning logs flag string
	WARNING string = "WARNING"
	//INFO informative log message flag string
	INFO string = "INFO"
	//DEBUG debug logs flag string
	DEBUG string = "DEBUG"
)

//LogLevels defines loglevel priorities  0 highest and 3 lowest
var LogLevels = map[string]int{MONITOR: 0, ERROR: 1, WARNING: 2, INFO: 3, DEBUG: 4}
