package logger

import "fmt"

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {

	fmt.Println("[LOG]", message)
}
