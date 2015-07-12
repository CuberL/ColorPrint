// ColorPrint project ColorPrint_linux.go
package ColorPrint

import (
	"fmt"
)

var (
	Black        = 0
	Blue         = 1
	Green        = 2
	Cyan         = 3
	Red          = 4
	Magenta      = 5
	Yellow       = 6
	White        = 7
	Gray         = 8
	LightBlue    = 9
	LightGreen   = 10
	LightCyan    = 11
	LightRed     = 12
	LightMagenta = 13
	Lightyellow  = 14
	LightWhite   = 15
)

var linux_color = []string{
	"\033[22;30m", // 0-black
	"\033[22;34m", // 1-blue
	"\033[22;32m", // 2-green
	"\033[22;36m", // 3-cyan
	"\033[22;31m", // 4-red
	"\033[22;35m", // 5-megenta
	"\033[22;33m", // 6-yellow
	"\033[22;37m", // 7-white
	"\033[01;30m", // 8-gray
	"\033[01;34m", // 9-lightblue
	"\033[01;32m", // 10-lightgreen
	"\033[01;36m", // 11-lightcyan
	"\033[01;31m", // 12-lightred
	"\033[01;35m", // 13-lightmagenta
	"\033[01;33m", // 14-lightyellow
	"\033[01;37m", // 15-lightwhite
}

func Print(color int, args ...interface{}) {
	if color < 0 || color >= 16 {
		fmt.Print(args...)
		return
	}
	fmt.Print(linux_color[color])
	fmt.Print(args...)
	fmt.Print("\033[0m")
}

func Printf(color int, format string, args ...interface{}) {
	if color < 0 || color >= 16 {
		fmt.Printf(format, args...)
		return
	}
	fmt.Print(linux_color[color])
	fmt.Printf(format, args...)
	fmt.Print("\033[0m")
}
