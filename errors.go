package helpers

import (
	"fmt"
	"os"
)

// ErrorCheck checks for err != nil and prints a message  in the format
// of "Error while <setting>: <err>" before exiting if it applies
func ErrorCheck(err error, setting string) {
	if err != nil {
		fmt.Printf("Error while %s: %s\n", setting, err)
		os.Exit(1)
	}
}

// Faile prints a message and calls os.Exit(1)
func Fail(args ...interface{}) {
	fmt.Println(args...)
	os.Exit(1)
}
