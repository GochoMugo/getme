package lib

import (
	"fmt"
	"github.com/mitchellh/colorstring"
)

const padding = "    "

func log(message string, color string) {
	colorstring.Println(padding + "⇒ [" + color + "]" + message)
}

func LogSuccess(message string) {
	log(message, "green")
}

func LogError(message string) {
	log(message, "red")
}

func LogNormal(message string) {
	log(message, "white")
}

func LogRaw(message string) {
	fmt.Println(message)
}
