package client

import (
	"io"
	"log"
	"net/http"
)

func httpClient(method, url, token string, body io.Reader) *http.Response {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatalf("Request: %v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", token)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Response: %v", err)
	}

	return response
}
