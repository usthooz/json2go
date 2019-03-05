package ozlog

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"sync"
)

// Logger is our golog.
type Logger struct {
	Prefix     string
	Level      Level
	TimeFormat string
	ShowLine   bool
	StackSkip  int
	Colorful   bool
	mu         sync.Mutex
	l          *log.Logger
}

// New info
func New() *Logger {
	return &Logger{
		Level:      InfoLevel,
		TimeFormat: "2006/01/02 15:04",
		ShowLine:   true,
		Colorful:   true,
		StackSkip:  4,
		l:          log.New(os.Stdout, "", log.LstdFlags),
	}
}

// SetPrefix
func (l *Logger) SetPrefix(prefix string) *Logger {
	l.Prefix = prefix
	return l
}

// Skip
func (l *Logger) Skip(skip int) *Logger {
	nl := *l
	nl.StackSkip = skip
	return &nl
}

// Returns itself.
func (l *Logger) SetTimeFormat(s string) *Logger {
	l.mu.Lock()
	l.TimeFormat = s
	l.mu.Unlock()
	return l
}

// Returns itself.
func (l *Logger) DisableNewLine() *Logger {
	l.mu.Lock()
	l.ShowLine = false
	l.mu.Unlock()
	return l
}

/*
	"disable"
	"fatal"
	"error"
	"warn"
	"info"
	"debug"
*/
// SetLevel
func (l *Logger) SetLevel(levelName string) *Logger {
	l.mu.Lock()
	l.Level = getLevelByName(levelName)
	l.mu.Unlock()
	return l
}

// getPrefix
func (l *Logger) getPrefix(level Level) string {
	var (
		s string
	)
	levels, ok := Levels[level]
	if !ok {
		return s
	}
	if l.Colorful {
		s = levels.ColorfulText
	}
	if l.Prefix != "" {
		s = s + " " + l.Prefix
	}
	return s
}

// getPosix
func (l *Logger) getPosix() string {
	if !l.ShowLine {
		return ""
	}
	_, file, line, ok := runtime.Caller(l.StackSkip)

	if !ok {
		return ""
	}
	file = path.Base(file)
	return fmt.Sprintf("[%s:%d]", file, line)
}

// output
func (l *Logger) output(level Level, format string, a ...interface{}) {
	if level > l.Level {
		return
	}
	var (
		s string
	)
	if len(a) != 0 {
		s = fmt.Sprintf(format, a...)
	} else {
		s = fmt.Sprint(format)
	}
	content := l.getPrefix(level) + " " + s + " " + l.getPosix()
	l.l.Println(content)
}

// Tracef
func (l *Logger) Tracef(format string, a ...interface{}) {
	l.output(TraceLevel, format, a...)
}

// Debugf
func (l *Logger) Debugf(format string, a ...interface{}) {
	l.output(DebugLevel, format, a...)
}

// Infof
func (l *Logger) Infof(format string, a ...interface{}) {
	l.output(InfoLevel, format, a...)
}

// Warnf
func (l *Logger) Warnf(format string, a ...interface{}) {
	l.output(WarnLevel, format, a...)
}

// Errorf
func (l *Logger) Errorf(format string, a ...interface{}) {
	l.output(ErrorLevel, format, a...)
}

// Fatalf
func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.output(FatalLevel, format, a...)
	os.Exit(1)
}
