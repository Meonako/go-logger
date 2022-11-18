package utils

import (
	"os"
	"runtime"
	"strings"
)

// See: runtime.Caller() for information
//
// Return: file, line, function name where depth called
func GetCodeInfo(depth int) (string, int, string) {
	pc, file, line, ok := runtime.Caller(depth)
	if ok {
		return getFileName(file), line, onlyFunctionName(runtime.FuncForPC(pc).Name())
	}

	return "Un-recoverable", -1, "Un-recoverable"
}

func getFileName(fullpath string) string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return fullpath[len(strings.ReplaceAll(dir, "\\", "/")+"/"):]
}

func onlyFunctionName(funcpath string) string {
	return LastElement(strings.Split(funcpath, "/"))
}
