package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/southofrubick/ikea-assignment/db"
	"github.com/southofrubick/ikea-assignment/entity"
)

func GetAllProducts(e *echo.Echo, pool *pgxpool.Pool) {
	e.GET("/api/products", func(c echo.Context) error {
		var products []entity.Product

		products, err := db.GetAllProducts(pool)
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
}
