package utils

import (
	"bufio"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

// Get user input. First argument is the message that will display on terminal before wait for user input.
//
// If you are familiar with Python's input(), this is basically the same thing.
func GetInput(msg ...string) string {
	if len(msg) > 0 {
		fmt.Print(msg[0])
	}

	scanner.Scan()
	return scanner.Text()
}
