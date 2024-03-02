package main

import (
	"fmt"
	"http-client-go/client"
	"http-client-go/model"
)

const (
	URL = "http://localhost:8080"
)

func main() {
	lc := client.LoginClient(URL+"/v1/login", "demo@gmail.com", "demo")
	product := model.Product{
		Name:         "Maiker",
		Observations: "Creating http client",
		Price:        20,
	}
	gr := client.ProductCreate(URL+"/v1/products", lc.Data.Token, &product)
	fmt.Println(gr)
}
