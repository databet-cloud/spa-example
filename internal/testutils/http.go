package testutils

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func MustMakeResponse(status int, filePath string) *http.Response {
	var body io.ReadCloser = http.NoBody

	if filePath != "" {
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}

		body = io.NopCloser(bytes.NewReader(fileContent))
	}

	return &http.Response{
		Status:     http.StatusText(status),
		StatusCode: status,
		Body:       body,
	}
}

func MustMakeRequest(method string, url string, body io.Reader) *http.Request {
	httpReq, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}

	return httpReq
}
