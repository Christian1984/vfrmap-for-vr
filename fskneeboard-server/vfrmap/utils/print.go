package utils

import (
	"fmt"
)

func Print(a ...interface{}) {
	fmt.Print(a...)

	//str := fmt.Sprint(a...)
	//controlpanel.ConsoleLog(str)
}

func Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)

	//str := fmt.Sprintf(format, a...)
	//controlpanel.ConsoleLog(str)
}

func Println(a ...interface{}) {
	fmt.Println(a...)

	//str := fmt.Sprintln(a...)
	//controlpanel.ConsoleLog(str)
}