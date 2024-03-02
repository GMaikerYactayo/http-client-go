package client

import (
	"bytes"
	"encoding/json"
	"http-client-go/model"
	"io"
	"log"
	"net/http"
)

func LoginClient(url, email, password string) model.LoginResponse {
	login := model.Login{
		Email:    email,
		Password: password,
	}

	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(&login)
	if err != nil {
		log.Fatalf("Error in marshal of login: %v", err)
	}

	response := httpClient(http.MethodPost, url, "", data)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading the body in login: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		log.Fatalf("A code 201 is expected, it was obtained: %d, answer: %s", response.StatusCode, string(body))
	}

	dataResponse := model.LoginResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("Error in the body marshal in login: %v", err)
	}

	return dataResponse
}
