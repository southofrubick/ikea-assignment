package API

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/southofrubick/ikea-assignment/db"
	"github.com/southofrubick/ikea-assignment/entity"
)

func GetAllProductTypes(e *echo.Echo, pool *pgxpool.Pool) {
	e.GET("/api/products/types", func(c echo.Context) error {
		var productTypes []entity.ProductType

		productTypes, err := db.GetAllProductTypes(pool)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to fetch productTypes")
		}
		log.Printf("Fetched %d product types from the database", len(productTypes))

		productTypesJson, err := json.Marshal(productTypes)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to marshal products to JSON")
		}

		return c.String(http.StatusOK, string(productTypesJson))
	})
}
