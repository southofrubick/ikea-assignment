package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

	e.Logger.Fatal(e.Start(":8080"))
}
