package logger

import (
	"fmt"
	"github.com/gookit/color"
	"time"
)

func PrintColor(c color.Color, mode string, value interface{}) {
	now := time.Now()
	message := fmt.Sprintf("[%s][%s] %v", now.Format("2006-01-02 15:04:05"), mode, value)
	c.Println(message)
}

func Success(value interface{}) {
	PrintColor(color.Green, "Success", value)
}

func Error(value interface{}) {
	PrintColor(color.Red, "Error", value)
}

func Warning(value interface{}) {
	PrintColor(color.Yellow, "Warning", value)
}

func Info(value interface{}) {
	PrintColor(color.Cyan, "Info", value)
}
