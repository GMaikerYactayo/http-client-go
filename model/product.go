package model

// Product model
type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Observations string `json:"observations"`
	Price        int    `json:"price"`
}
