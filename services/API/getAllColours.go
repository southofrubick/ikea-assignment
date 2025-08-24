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

func GetAllColours(e *echo.Echo, pool *pgxpool.Pool) {
	e.GET("/api/products/colours", func(c echo.Context) error {
		var colours []entity.Colour

		colours, err := db.GetAllColours(pool)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to fetch colours")
		}
		log.Printf("Fetched %d colours from the database", len(colours))

		coloursJson, err := json.Marshal(colours)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to marshal products to JSON")
		}

		return c.String(http.StatusOK, string(coloursJson))
	})
}
