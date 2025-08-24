package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/southofrubick/ikea-assignment/db"
	"github.com/southofrubick/ikea-assignment/entity"
)


func main() {
	pool, err := db.InitDB()

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer pool.Close()

	e := echo.New()

	e.GET("/api/products", func(c echo.Context) error {
		var products []entity.Product

		products, err = db.GetAllProducts(pool)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to fetch products")
		}
		log.Printf("Fetched %d products from the database", len(products))

		productsJson, err := json.Marshal(products)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to marshal products to JSON")
		}

		return c.String(http.StatusOK, string(productsJson))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
