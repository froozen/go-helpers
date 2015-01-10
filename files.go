package helpers

import (
	"io/ioutil"
	"os"
)

// ReadFile opens a file and returns its contents
func ReadFile(filename string) (string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the whole file
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(all), nil
}
