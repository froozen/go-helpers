package helpers

import (
	"io/ioutil"
	"net/http"
)

// RetrievePage wraps the whole process of getting a websites code
// from the internet. The equivalent of pythons urllib.retrieve()
func RetrievePage(url string) (string, error) {
	// Get the page
	page, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer page.Body.Close()

	// Read the whole page
	content, err := ioutil.ReadAll(page.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
