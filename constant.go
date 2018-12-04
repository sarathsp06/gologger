package logger

const (
	// MONITOR monitoring logs flag
	MONITOR string = "MON"
	// ERROR error logs flag
	ERROR string = "ERR"
	// WARNING Warning logs flag string
	WARNING string = "WAR"
	// INFO informative log message flag string
	INFO string = "INF"
	// DEBUG debug logs flag string
	DEBUG string = "DEB"
)

// Constants representing colors to tty color constants
// to be used as "\033[COLORm"
// Eg: \033[31m for RED
const (
	WHITE     = 0
	RED   int = iota + 31
	GREEN
	YELLOW
	BLUE
	MAGENTA
)

//LogLevels defines loglevel priorities  0 highest and 3 lowest
var LogLevels = map[string]int{MONITOR: 0, ERROR: 1, WARNING: 2, INFO: 3, DEBUG: 4}

// LogColors is map of log level to color
var LogColors = map[string]int{
	MONITOR: MAGENTA,
	ERROR:   RED,
	WARNING: GREEN,
	INFO:    BLUE,
	DEBUG:   YELLOW,
}
