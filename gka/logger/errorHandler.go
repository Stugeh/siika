package logger

import "fmt"

func error(lineNumber int, message string) {
	report(lineNumber, "", message)
}

func report(lineNumber int, location string, message string) {
	fmt.Println(fmt.Sprintf("[line: %d] Error %s: %s", lineNumber, location, message))
}
