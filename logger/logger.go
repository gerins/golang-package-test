package logger

import "fmt"

var globalVar string

func Init(plugins string) {
	globalVar = plugins
}

func Info(value string) {
	if globalVar == "" {
		panic("neet to init first")
	}

	fmt.Println(globalVar + value)
}
