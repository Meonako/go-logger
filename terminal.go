package logger

import (
	"os"
	"strconv"
	"time"

	"github.com/Meonako/go-logger/utils"

	"github.com/fatih/color"
)

var (
	logger = GetDefault()
)

func init() {
	logger.UpdatePrefix = updatePrefix
}

func updatePrefix() string {
	file, line, function := utils.GetCodeInfo(4)
	return Magenta(time.Now().Format("02/01/2006 15:04:05 ")) +
		Colorize(file+":"+strconv.Itoa(line)+" | "+function, color.BgCyan, color.FgBlack) +
		" -> "
}

// Print to terminal with new line at the end. Equivalent to fmt.Println.
func ToTerminal(args ...any) {
	logger.Println(args...)
}

// Print to terminal with format string. Equivalent to fmt.Printf with \n at the end.
func ToTerminalFormat(format string, v ...any) {
	logger.Printf(format, v...)
}

// Print to terminal. Exit with code 0 afterwards.
func ToTerminalAndExit(args ...any) {
	logger.Println(args...)
	utils.GetInput("Press any key to exit...")
	os.Exit(0)
}

// Print to terminal with format string and \n at the end. Exit with code 0 afterwards.
func ToTerminalFormatAndExit(format string, v ...any) {
	logger.Printf(format, v...)
	utils.GetInput("Press any key to exit...")
	os.Exit(0)
}

// Print to terminal without new line at the end. Equivalent to fmt.Print.
func ToTerminalWithoutNewLine(args ...any) {
	logger.Print(args...)
}

// Print to terminal if error (first argument) is not nil.
// If args (second+ arguments) is NOT passed,
// It'll print the error (first argument) message.
func ToTerminalIfError(err error, args ...any) {
	if err == nil {
		return
	}

	if len(args) > 0 {
		logger.Println(args...)
		return
	}

	logger.Println(err)
}

// Print to termianl with format string if error (first argument) is not nil.
// if v is NOT passed,
// It'll print the format with err.
//
// like this -> fmt.Printf(format, err)
func ToTerminalFormatIfError(err error, format string, v ...any) {
	if err == nil {
		return
	}

	if len(v) > 0 {
		logger.Printf(format, v...)
		return
	}

	logger.Printf(format, err)
}

// Print to terminal if error (first argument) is not nil.
// If args (second+ arguments) is NOT passed,
// It'll print the error (first argument) message.
//
// Exit if error (first argument)
func ToTerminalAndExitIfError(err error, args ...any) {
	if err == nil {
		return
	}

	if len(args) > 0 {
		logger.Println(args...)
	} else {
		logger.Println(err)
	}

	utils.GetInput("Press any key to exit...")
	os.Exit(0)
}

// Print to termianl with format string if error (first argument) is not nil.
// if v is NOT passed,
// It'll print the format with err.
//
// like this -> fmt.Printf(format, err)
//
// Exit if error (first argument)
func ToTerminalAndExitFormatIfError(err error, format string, v ...any) {
	if err == nil {
		return
	}

	if len(v) > 0 {
		logger.Printf(format, v...)
	} else {
		logger.Printf(format, err)
	}

	utils.GetInput("Press any key to exit...")
	os.Exit(0)
}

// --------------------------------------------------
// ---------------------  RED  ----------------------
// --------------------------------------------------

// Print to terminal with red text and new line at the end.
func ToTerminalRed(args ...any) {
	newFunc := GetColorizeFunc(color.FgRed)
	logger.Println(newFunc(args...))
}

// Print to terminal with red text formatted string and new line at the end.
func ToTerminalRedFormat(format string, v ...any) {
	logger.Printf(Red(format), v...)
}

// Print to terminal with red text. Exit with code 0 afterwards.
func ToTerminalRedAndExit(args ...any) {
	newFunc := GetColorizeFunc(color.FgRed)
	logger.Println(newFunc(args...))
	utils.GetInput("Press any key to exit...")
	os.Exit(0)
}

// Print to terminal with red formatted string and \n at the end. Exit with code 0 afterwards.
func ToTerminalRedFormatAndExit(format string, v ...any) {
	logger.Printf(Red(format), v...)
	utils.GetInput("Press any key to exit...")
	os.Exit(0)
}

// Print to terminal with red text if error (first argument) is not nil.
// If args (second+ arguments) is NOT passed,
// It'll print the error (first argument) message.
func ToTerminalRedIfError(err error, args ...any) {
	if err == nil {
		return
	}

	newFunc := GetColorizeFunc(color.FgRed)

	if len(args) > 0 {
		logger.Println(newFunc(args...))
		return
	}

	logger.Println(newFunc(err))
}

// Print to terminal with formatted red text if error (first argument) is not nil.
// If v (second+ arguments) is NOT passed,
// It'll print the format with err.
//
// like this -> fmt.Printf(format, err)
func ToTerminalRedFormatIfError(err error, format string, v ...any) {
	if err == nil {
		return
	}

	if len(v) > 0 {
		logger.Printf(Red(format), v...)
		return
	}

	logger.Printf(Red(format), err)
}

// Print to terminal if error (first argument) is not nil.
// If args (second+ arguments) is NOT passed,
// It'll print the error (first argument) message.
//
// Exit if error (first argument)
func ToTerminalRedAndExitIfError(err error, args ...any) {
	if err == nil {
		return
	}

	newFunc := GetColorizeFunc(color.FgRed)

	if len(args) > 0 {
		logger.Println(newFunc(args...))
	} else {
		logger.Println(newFunc(err))
	}

	utils.GetInput("Press any key to exit...")
	os.Exit(0)
}

// Print to termianl with format string if error (first argument) is not nil.
// if v is NOT passed,
// It'll print the format with err.
//
// like this -> fmt.Printf(format, err)
//
// Exit if error (first argument)
func ToTerminalRedAndExitFormatIfError(err error, format string, v ...any) {
	if err == nil {
		return
	}

	if len(v) > 0 {
		logger.Printf(Red(format), v...)
	} else {
		logger.Printf(Red(format), err)
	}

	utils.GetInput("Press any key to exit...")
	os.Exit(0)
}
