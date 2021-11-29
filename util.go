package CurseAPIWrapper

import (
	"bytes"
	"errors"
	"net/http"
)

func makeRequest(method, url string, b []byte) (*http.Response, error) {
	if apiKey == "" {
		return nil, errors.New("API key not set")
    }
	headers := map[string][]string{
		"Accept": []string{"application/json"},
		"Content-Type": []string{"application/json"},
		"x-api-key": []string{apiKey},
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header = headers

	return client.Do(req)
}