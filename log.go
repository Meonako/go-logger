package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/Meonako/go-logger/utils"
)

// All var and const should start with lower case because we don't want to export to other pkg

var (
	prefix  string
	logFile *os.File

	errCount = 0
)

// Init function for file logging.
func Init() {
	if !Settings.LogToFile {
		return
	}

	var err error
	if _, err = os.Stat(Settings.LogFolder); os.IsNotExist(err) {
		err = os.Mkdir(Settings.LogFolder, 0700)
		if err != nil {
			fmt.Println("LOGGER : Can not create dir \"" + Settings.LogFolder + "\". logger module disabled")
			writeErrorFile("LOGGER : Can not create dir \"" + Settings.LogFolder + "\". logger module disabled")
			errCount++
		}
	}

	logFile, err = os.Create(Settings.LogFolder + "/" + time.Now().Format(Settings.LogFileName))
	if err != nil {
		fmt.Println("LOGGER : Can not create log file. logger module disabled")
		writeErrorFile("LOGGER : Can not create log file. logger module disabled")
		errCount++
	}
}

// ------------------------------- LOCAL FUNCTIONS -------------------------------
// Local functions, as the section name said, are the function that only available
// to this pkg only. ALWAYS start with lower case

func updatePrefix(newPrefix string) {
	if errCount != 0 {
		return
	}

	prefix = time.Now().Format(Settings.DateFormat+" ") + newPrefix + " "
}

func writeLogFile(text string) {
	logFile.Write([]byte(text + "\n"))
}

func writeErrorFile(text string) {
	file, err := os.Create("error.log")
	if err != nil {
		return
	}

	_, err = file.Write([]byte(text))
	if err != nil {
		return
	}
}

func println(text string) {
	fmt.Println(prefix + text)

	if errCount != 0 || !Settings.LogToFile {
		return
	}
	writeLogFile(prefix + text)
}

func printf(format string, v ...any) {
	fmt.Printf(prefix+format+"\n", v...)

	if errCount != 0 || !Settings.LogToFile {
		return
	}
	writeLogFile(fmt.Sprintf(prefix+format, v...))
}

func exit(code ...int) {
	utils.GetInput("Press any key to exit...")
	if len(code) > 0 {
		os.Exit(code[0])
	}

	os.Exit(0)
}

// ------------------------------- GLOBAL FUNCTIONS -------------------------------
// Global functions are the function that other pkg can call.
// ALWAYS start with upper case

// Log info level.
func Info(text string) {
	updatePrefix(Settings.InfoPrefix)
	println(text)
}

// Log info level with format string.
func Infof(format string, v ...any) {
	updatePrefix(Settings.InfoPrefix)
	printf(format, v...)
}

// Same as 'Info' but exit the program afterward. This will prompt to press any key to exit.
func InfoAndExit(text string) {
	Info(text)
	exit()
}

// Same as 'Infof' but exit afterward. This will prompt to press any key to exit.
func InfofAndExit(format string, v ...any) {
	Infof(format, v...)
	exit()
}

// Log warn level is for when error occured but you DONT want to stop your program (i.e. the error that can be ignore)
func Warn(text string) {
	updatePrefix(Settings.WarnPrefix)
	println(text)
}

// Log warn level with format string. See 'logger.Warn()' for more information.
func Warnf(format string, v ...any) {
	updatePrefix(Settings.WarnPrefix)
	printf(format, v...)
}

// This function is useful when you're lazy to write 'if err != nil { logger.Warn(err.Error()) }'.
//
// 'text' argument is optional, in case you don't want to 'logger.Warn(err.Error())' but to be 'logger.Warn("custom")' instead. ONLY first one will be used.
func WarnIf(err error, text ...string) {
	if err == nil {
		return
	}

	if len(text) > 0 {
		Warn(text[0])
		return
	}

	Warn(err.Error())
}

// This function is useful when you're lazy to write 'if err != nil { logger.Warnf("custom format : %v", err) }'.
//
// 'v' argument is optional. If nothing is passed, it will use 'err' as first and only argument.
func WarnIff(err error, format string, v ...any) {
	if err == nil {
		return
	}

	if len(v) > 0 {
		Warnf(format, v...)
		return
	}

	Warnf(format, err)
}

// Log error level is for when error occured and you NEED TO STOP your program. This will prompt to press any key to exit.
func Error(text string) {
	updatePrefix(Settings.ErrorPrefix)
	println(text)
	exit(1)
}

// Log error level with format string. See 'logger.Error()' for more information.
func Errorf(format string, v ...any) {
	updatePrefix(Settings.ErrorPrefix)
	printf(format, v...)
	exit(1)
}

// This function is useful when you're lazy to write 'if err != nil { logger.Error(err.Error()) }'.
//
// THIS WILL STOP YOUR PROGRAM. See 'logger.WarnIf()' if you don't want to.
//
// 'text' argument is optional, in case you don't want to 'logger.Error(err.Error())' but to be 'logger.Error("custom")' instead. ONLY first one will be used.
func ErrorIf(err error, text ...string) {
	if err == nil {
		return
	}

	if len(text) > 0 {
		Error(text[0])
	}

	Error(err.Error())
}

// This function is useful when you're lazy to write 'if err != nil { logger.Errorf("custom format : %v", err) }'.
//
// THIS WILL STOP YOUR PROGRAM. See 'logger.WarnIff()' if you don't want to.
//
// 'v' argument is optional. If nothing is passed, it will use 'err' as first and only argument.
func ErrorIff(err error, format string, v ...any) {
	if err == nil {
		return
	}

	if len(v) > 0 {
		Errorf(format, v...)
	}

	Errorf(format, err)
}
