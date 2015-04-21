package lib

import (
	"fmt"
	"github.com/mitchellh/colorstring"
	"github.com/ttacon/chalk"
	"strings"
)

const padding = "    "

func log(message []string, color string) {
	colorstring.Println(padding + "⇒ [" + color + "]" + strings.Join(message, ""))
}

func LogSuccess(message ...string) {
	log(message, "green")
}

func LogError(message ...string) {
	log(message, "red")
}

func LogNormal(message ...string) {
	log(message, "white")
}

func LogRaw(message ...string) {
	fmt.Println(message)
}

func TextBold(str string) string {
  return chalk.Bold.TextStyle(str)
}

func TextUnderline(str string) string {
  return chalk.Underline.TextStyle(str)
}

