package mwgoapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Handler is the handler for the API.
type Handler struct {
	Client  *http.Client
	BaseURL string
	Key     string
}

const (
	// BaseURL is the base URL for the API.
	BaseURL      = "https://www.dictionaryapi.com/api/v3/references/collegiate/json"
	AudioFormat  = "mp3"
	AudioBaseURL = "https://media.merriam-webster.com/audio/prons/en/us"
)

// NewClient returns a new Handler.
func NewClient(client *http.Client, url, key string) *Handler {
	if url == "" {
		url = BaseURL
	}
	return &Handler{Client: client, BaseURL: url, Key: key}
}

// Get returns the response for the given word.
func (a *Handler) Get(word string) ([]byte, error) {
	res, err := a.Client.Get(fmt.Sprintf("%s/%s?key=%s", a.BaseURL, word, a.Key))
	if err != nil {
		return nil, fmt.Errorf("could not send request, %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read body, err: %w", err)
	}

	if string(b) == "Invalid API key. Not subscribed for this reference." {
		return nil, fmt.Errorf("invalid API key")
	}

	return b, nil
}
