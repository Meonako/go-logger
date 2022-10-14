package logger

import (
	"strings"
)

type settings struct {
	// Set to false if you don't want to log to file. Default is true
	LogToFile bool
	// Log folder name
	LogFolder,
	// Log file name
	LogFileName string

	// Info-level prefix
	InfoPrefix,
	// Warn-level prefix
	WarnPrefix,
	// Error-leve prefix
	ErrorPrefix string
}

// Settings for logger module with default value
var Settings = settings{
	LogToFile:   true,
	LogFolder:   "log",
	LogFileName: "02_01_2006 15-04-05.log",

	InfoPrefix:  "[ INFO ]:",
	WarnPrefix:  "[ WARN ]:",
	ErrorPrefix: "[ ERROR ]:",
}

var humanDateFormat = []string{
	"dd",
	"mm",
	"yyyy",
	"hh",
	"mm",
	"ss",
}

var goDateFormat = []string{
	"02",
	"01",
	"2006",
	"15",
	"04",
	"05",
}

// You can actually set these properties directly. But this is for one-liner lover. Pass "-" (hyphen or dash) if you don't want to change value.
//
// For fileName, people usually use time-date format for log's filename. I made it a lil bit easier.
// You can use "mm-dd-yyyy" instead of that "01-02-2006".
//
// It will replace one by one so for example "yyyy-mm-dd hh!mm!ss.log" will result in "2006-01-02 15!04!05.log" and then time.Format will take care of it!
func (s *settings) Set(folderName, fileName, infoPrefix, warnPrefix, errorPrefix string) {
	if fileName != "-" {
		replace(&fileName, humanDateFormat, goDateFormat)
	}

	Settings.LogFolder = checkDefault(folderName, Settings.LogFolder)
	Settings.LogFileName = checkDefault(fileName, Settings.LogFileName)

	Settings.InfoPrefix = checkDefault(infoPrefix, Settings.InfoPrefix)
	Settings.WarnPrefix = checkDefault(warnPrefix, Settings.WarnPrefix)
	Settings.ErrorPrefix = checkDefault(errorPrefix, Settings.ErrorPrefix)
}

func replace(str *string, old []string, new []string) {
	if len(old) != len(new) {
		return
	}

	for i := 0; i < len(old); i++ {
		result := strings.ReplaceAll(*str, old[i], new[i])
		str = &result
	}
}

func checkDefault(str string, def string) string {
	if str == "-" {
		return def
	}

	return str
}
