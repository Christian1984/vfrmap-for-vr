package utils

import (
	"fmt"
)

var guiPrintCallback func(string)

func SetGuiPrintCallback(callback func(string)) {
	guiPrintCallback = callback
}

func printToGui(message string) {
	if guiPrintCallback != nil {
		go guiPrintCallback(message)
	}
}

func Print(a ...interface{}) {
	fmt.Print(a...)

	str := fmt.Sprint(a...)
	printToGui(str)
}

func Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)

	str := fmt.Sprintf(format, a...)
	printToGui(str)
}

func Println(a ...interface{}) {
	fmt.Println(a...)

	str := fmt.Sprintln(a...)
	printToGui(str)
}