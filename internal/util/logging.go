package util

import "fmt"

func LogCon(s string, a any, logType int) {
	switch logType {
	case Normal:
		fmt.Println(s)
	case Error:
		fmt.Println("[E]: ", s, a)
	case Warning:
		fmt.Println("[W]: ", s, a)
	case Info:
		fmt.Println("[I]: ", s, a)
	}
}

// enum with logging types, makes calling LogCon easier
const (
	Normal = iota
	Error
	Warning
	Info
)
