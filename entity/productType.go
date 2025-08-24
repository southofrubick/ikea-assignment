package entity

import "time"

type ProductType struct {
	ID   int `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
