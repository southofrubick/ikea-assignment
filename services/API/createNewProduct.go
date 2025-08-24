package API

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/southofrubick/ikea-assignment/db"
)

func CreateNewProduct(e *echo.Echo, pool *pgxpool.Pool) {
	e.POST("/api/products", func(c echo.Context) error {
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			log.Println("Failed to decode JSON body", err)
			return c.String(http.StatusBadRequest, "Invalid JSON body")
		}

		name := json_map["name"].(string)
		product_type_id, err := strconv.Atoi(json_map["product_type_id"].(string))
		if err != nil {
			log.Println("Invalid product_type_id", err)
			return c.String(http.StatusBadRequest, "Invalid product_type_id")
		}
		colourID, err := strconv.Atoi(json_map["colour_id"].(string))
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
