package API

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/southofrubick/ikea-assignment/db"
)

func CreateNewProduct(e *echo.Echo, pool *pgxpool.Pool) {
	e.POST("/api/products", func(c echo.Context) error {
		name := c.FormValue("name")
		product_type_id, err := strconv.Atoi(c.FormValue("product_type_id"))
		if err != nil {
			log.Println("Invalid product_type_id", err)
			return c.String(http.StatusBadRequest, "Invalid product_type_id")
		}
		colourID, err := strconv.Atoi(c.FormValue("colour_id"))
		if err != nil {
			log.Println("Invalid colour_id", err)
			return c.String(http.StatusBadRequest, "Invalid colour_id")
		}
		if name == "" {
			log.Println("Name is required")
			return c.String(http.StatusBadRequest, "Name is required")
		}

		productID, err := db.CreateProduct(pool, name, product_type_id, colourID)
		if err != nil {
			log.Println("Failed to create product", err)
			return c.String(http.StatusInternalServerError, "Failed to create product")
		}

		log.Printf("Created product with ID %d", productID)

		return c.String(http.StatusCreated, string(productID))
	})
}
