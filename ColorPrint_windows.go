// ColorPrint project ColorPrint_windows.go
package ColorPrint

import (
	"fmt"
	"syscall"
	"unsafe"
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

type (
	short int16

	word uint16

	coord struct {
		x short
		y short
	}

	small_rect struct {
		left   short
		top    short
		right  short
		bottom short
	}

	console_screen_buffer_info struct {
		size                coord
		cursor_position     coord
		attributes          int64
		window              small_rect
		maximum_window_size coord
	}
)

var tmp console_screen_buffer_info
var kernel32 syscall.Handle
var hnd, GetStdHandle, SetConsoleTextAttribute, GetConsoleScreenBufferInfo uintptr

func init() {
	var nStdHandle = -11
	kernel32, _ = syscall.LoadLibrary("Kernel32.dll")
	GetStdHandle, _ = syscall.GetProcAddress(syscall.Handle(kernel32), "GetStdHandle")
	SetConsoleTextAttribute, _ = syscall.GetProcAddress(syscall.Handle(kernel32), "SetConsoleTextAttribute")
	//	ReadConsoleOutputAttribute, _ := syscall.GetProcAddress(syscall.Handle(kernel32), "ReadConsoleOutputAttribute")
	GetConsoleScreenBufferInfo, _ = syscall.GetProcAddress(syscall.Handle(kernel32), "GetConsoleScreenBufferInfo")
	hnd, _, _ = syscall.Syscall(uintptr(GetStdHandle), 1, uintptr(nStdHandle), 0, 0)
}

func Print(color int, args ...interface{}) {
	if color < 0 || color >= 16 {
		fmt.Print(args...)
		return
	}
	syscall.Syscall(uintptr(GetConsoleScreenBufferInfo), 2, hnd, uintptr(unsafe.Pointer(&tmp)), 0)
	syscall.Syscall(uintptr(SetConsoleTextAttribute), 2, hnd, uintptr(color), 0)
	fmt.Print(args...)
	syscall.Syscall(uintptr(SetConsoleTextAttribute), 2, hnd, uintptr(tmp.attributes), 0)
}

func Printf(color int, format string, args ...interface{}) {
	if color < 0 || color >= 16 {
		fmt.Printf(format, args...)
		return
	}
	syscall.Syscall(uintptr(GetConsoleScreenBufferInfo), 2, hnd, uintptr(unsafe.Pointer(&tmp)), 0)
	syscall.Syscall(uintptr(SetConsoleTextAttribute), 2, hnd, uintptr(color), 0)
	fmt.Printf(format, args...)
	syscall.Syscall(uintptr(SetConsoleTextAttribute), 2, hnd, uintptr(tmp.attributes), 0)
}
