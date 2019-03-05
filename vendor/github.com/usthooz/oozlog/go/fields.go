package ozlog

import (
	"time"
)

// A Log represents a log line.
type Log struct {
	// Logger
	Logger *Logger
	// Time
	Time time.Time
	// Level
	Level   Level
	Message string
	NewLine bool
}

// FormatTime current time
func (l Log) FormatTime() string {
	if l.Logger.TimeFormat == "" {
		return ""
	}
	return l.Time.Format(l.Logger.TimeFormat)
}
