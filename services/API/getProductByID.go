package API

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/southofrubick/ikea-assignment/db"
	"github.com/southofrubick/ikea-assignment/entity"
)

func GetProductByID(e *echo.Echo, pool *pgxpool.Pool) {
	e.GET("/api/products/:id", func(c echo.Context) error {
		var product *entity.Product

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid product ID")
		}

		product, err = db.GetProductByID(pool, id)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to fetch product")
		}
		log.Printf("Fetched %d product from the database", product)

		productJson, err := json.Marshal(product)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to marshal product to JSON")
		}

		return c.String(http.StatusOK, string(productJson))
	})
}
