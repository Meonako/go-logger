package logger

import (
	"github.com/fatih/color"
)

// Custom styles. See: https://github.com/fatih/color
func Colorize(text any, atr ...color.Attribute) string {
	Func := color.New(atr...).SprintFunc()
	return Func(text)
}

func GetColorizeFunc(attr ...color.Attribute) func(...any) string {
	return color.New(attr...).SprintFunc()
}

// Preset green for text.
func Green(text any) string {
	return Colorize(text, color.FgGreen)
}

// Preset yellow for text.
func Yellow(text any) string {
	return Colorize(text, color.FgYellow)
}

// Preset red for text.
func Red(text any) string {
	return Colorize(text, color.FgRed)
}

// Preset blue for text.
func Blue(text any) string {
	return Colorize(text, color.FgBlue)
}

// Preset cyan for text.
func Cyan(text any) string {
	return Colorize(text, color.FgCyan)
}

// Preset magenta for text.
func Magenta(text any) string {
	return Colorize(text, color.FgMagenta)
}
