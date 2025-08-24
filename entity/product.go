package entity

import "time"

type Product struct {
	ID   int `json:"id"`
	Name string `json:"name"`
	Colours []string `json:"colours,omitempty"`
	ProductTypeID int `json:"product_type_id"`
	ProductType string `json:"product_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
