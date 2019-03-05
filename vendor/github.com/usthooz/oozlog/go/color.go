package ozlog

import (
	"fmt"
)

var colorFormat = "\x1b[%dm%s\x1b[0m"

func colorize(colorCode int, str string) string {
	return fmt.Sprintf(colorFormat, colorCode, str)
}

// colorBlue code 蓝色
var colorBlue = 34

// blue color
func blue(s string) string {
	return colorize(colorBlue, s)
}

// colorLightGreen code 天蓝色
var colorLightGreen = 36

// lightGreen color
func lightGreen(s string) string {
	return colorize(colorLightGreen, s)
}

// colorPurple code 淡红
var colorPurple = 35

// purple color
func purple(s string) string {
	return colorize(colorPurple, s)
}

// colorRed code 红色
var colorRed = 31

// red color
func red(s string) string {
	return colorize(colorRed, s)
}

// colorGreen code 绿色
var colorGreen = 32

// green color
func green(s string) string {
	return colorize(colorGreen, s)
}

// colorYellow code 黄色
var colorYellow = 33

// yellow color
func yellow(s string) string {
	return colorize(colorYellow, s)
}

//  淡橙色
var colorOrange = 91

// orange color
func orange(s string) string {
	return colorize(colorOrange, s)
}
