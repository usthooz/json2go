package ozlog

// level 默认级别"info"
var Default = New()

// SetPrefix
func SetPrefix(s string) *Logger {
	return Default.SetPrefix(s)
}

// SetTimeFormat 时间格式设置
func SetTimeFormat(s string) {
	Default.SetTimeFormat(s)
}

/*
	"disable"
	"fatal"
	"error"
	"warn"
	"info"
	"debug"
*/
// SetLevel 设置输出级别
func SetLevel(levelName string) {
	Default.SetLevel(levelName)
}

// Fatalf
func Fatalf(format string, args ...interface{}) {
	Default.Fatalf(format, args...)
}

// Errorf
func Errorf(format string, args ...interface{}) {
	Default.Errorf(format, args...)
}

// Warnf
func Warnf(format string, args ...interface{}) {
	Default.Warnf(format, args...)
}

// Infof
func Infof(format string, args ...interface{}) {
	Default.Infof(format, args...)
}

// Debugf
func Debugf(format string, args ...interface{}) {
	Default.Debugf(format, args...)
}

// Tracef
func Tracef(format string, args ...interface{}) {
	Default.Tracef(format, args...)
}
