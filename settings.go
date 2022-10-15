package logger

import (
	"strings"
)

type Config struct {
	// Set to false if you don't want to log to file.
	LogToFile bool
	// Log folder name.
	LogFolder string
	// Log file name
	LogFileName string

	// Date format when logging
	DateFormat string

	// Info-level prefix
	InfoPrefix string
	// Warn-level prefix
	WarnPrefix string
	// Error-leve prefix
	ErrorPrefix string
}

// Settings for logger module with default value
var Settings = Config{
	LogToFile:   true,
	LogFolder:   "log",
	LogFileName: "02_01_2006 15-04-05.log",

	DateFormat: "02/01/2006 15:04:05",

	InfoPrefix:  "[ INFO ]:",
	WarnPrefix:  "[ WARN ]:",
	ErrorPrefix: "[ ERROR ]:",
}

var humanDateFormat = []string{
	"dd",   // 02
	"mm",   // 01
	"yyyy", // 2006
	"hh",   // 15
	"mm",   // 04
	"ss",   // 05
}

var goDateFormat = []string{
	"02",   // dd
	"01",   // mm
	"2006", // yyyy
	"15",   // hh
	"04",   // mm
	"05",   // ss
}

// Set settings by config struct. Pass "-" (hyphen or dash) if you don't want to change value.
//
// For LogFileName and DateFormat, people usually use time-date format for log's filename. I made it a lil bit easier.
// You can use "mm-dd-yyyy" instead of that "01-02-2006".
//
// It will replace one by one so for example "yyyy-mm-dd hh!mm!ss.log" will result in "2006-01-02 15!04!05.log" and then time.Format will take care of it!
func NewSettings(cfg Config) {
	if cfg.LogFileName != "-" {
		replace(&cfg.LogFileName, humanDateFormat, goDateFormat)
	}

	if cfg.DateFormat != "-" {
		replace(&cfg.DateFormat, humanDateFormat, goDateFormat)
	}

	Settings.LogToFile = cfg.LogToFile
	Settings.LogFolder = checkDefault(cfg.LogFolder, Settings.LogFolder)
	Settings.LogFileName = checkDefault(cfg.LogFileName, Settings.LogFileName)

	Settings.DateFormat = checkDefault(cfg.DateFormat, Settings.DateFormat)

	Settings.InfoPrefix = checkDefault(cfg.InfoPrefix, Settings.InfoPrefix)
	Settings.WarnPrefix = checkDefault(cfg.WarnPrefix, Settings.WarnPrefix)
	Settings.ErrorPrefix = checkDefault(cfg.ErrorPrefix, Settings.ErrorPrefix)
}

// You can actually set these properties directly or NewSettings. But this is for one-liner lover. Pass "-" (hyphen or dash) if you don't want to change value.
//
// For fileName and dateFormat, people usually use time-date format for log's filename. I made it a lil bit easier.
// You can use "mm-dd-yyyy" instead of that "01-02-2006".
//
// It will replace one by one so for example "yyyy-mm-dd hh!mm!ss.log" will result in "2006-01-02 15!04!05.log" and then time.Format will take care of it!
//
// 'Prefix' argument is optional, first arg is "InfoPrefix", second arg is "WarnPrefix", third arg is "ErrorPrefix". If not passed, use default.
func (s *Config) Set(logToFile bool, folderName, fileName, dateFormat string, Prefix ...string) {
	if fileName != "-" {
		replace(&fileName, humanDateFormat, goDateFormat)
	}

	Settings.LogToFile = logToFile
	Settings.LogFolder = checkDefault(folderName, Settings.LogFolder)
	Settings.LogFileName = checkDefault(fileName, Settings.LogFileName)

	Settings.DateFormat = checkDefault(dateFormat, Settings.DateFormat)

	Settings.InfoPrefix = Prefix[0]
	Settings.WarnPrefix = Prefix[1]
	Settings.ErrorPrefix = Prefix[2]
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
