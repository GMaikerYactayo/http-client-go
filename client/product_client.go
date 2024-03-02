package client

import (
	"bytes"
	"encoding/json"
	"http-client-go/model"
	"io"
	"log"
	"net/http"
)

func ProductCreate(url, token string, product *model.Product) model.GeneralResponse {
	data := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(data).Encode(product)
	if err != nil {
		log.Fatalf("Error in marshal of product: %v", err)
	}

	response := httpClient(http.MethodPost, url, token, data)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading the body in product: %v", err)
	}

	if response.StatusCode != http.StatusCreated {
		log.Fatalf("A code 201 is expected, it was obtained: %d, answer: %s", response.StatusCode, string(body))
	}

	dataResponse := model.GeneralResponse{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&dataResponse)
	if err != nil {
		log.Fatalf("Error in the body marshal in product: %v", err)
	}
	return dataResponse
}
