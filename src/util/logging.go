package util

import (
	"github.com/fatih/color"
)

func Log(message string) {
	color.White("["+ GetFormattedTime() + "] " + message)
}

func LogInfo(message string) {
	color.Blue("["+ GetFormattedTime() + "] " + message)
}

func LogSuccess(message string) {
	color.Green("["+ GetFormattedTime() + "] " + message)
}

func LogFail(message string) {
	color.Red("["+ GetFormattedTime() + "] " + message)
}

func LogError(message error) {
	color.Red("["+ GetFormattedTime() + "] " + message.Error())
}
