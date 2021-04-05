package logoutput

import (
	"fmt"
	"strings"
)

func ParseMessageByLevel(i int, log string, loglevel *int) {
	if *loglevel == 0 {
		PrintLog(i, log)
	}
	if *loglevel == 1 && HasError(log) {
		PrintLog(i, log)
	}
}

func PrintLog(i int, log string) {
	fmt.Printf("LOG %d:%s \n", i, log)
	fmt.Print(log)
}

func HasError(log string) bool {
	return strings.Contains(strings.ToLower(log), "stacktrace") || strings.Contains(strings.ToLower(log), "error")
}
