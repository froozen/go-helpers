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

// IsFile checks wether name corresponds to a file
func IsFile(name string) bool {
	_, err := os.Open(name)
	return err == nil
}

// IsDir checks wether name corresponds to a directory
func IsDir(name string) bool {
	dir, err := os.Stat(name)
	if err != nil {
		return false
	}
	return dir.IsDir()
}
