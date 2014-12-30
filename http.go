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

// RetrievePageCookies does the same as RetrievePage, but adds cookies to
// the http request.
func RetrievePageCookies(url string, cookies []*http.Cookie) (string, error) {
	// Crete the client
	client := &http.Client{}

	// Create the request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Add the cookies
	for _, cookie := range cookies {
		request.AddCookie(cookie)
	}

	// Send the request
	page, err := client.Do(request)
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
