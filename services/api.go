package services

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/southofrubick/ikea-assignment/services/API"
)

func InitAPI(pool *pgxpool.Pool) error {
	e := echo.New()
	e.Use(
		middleware.CORSWithConfig(
			middleware.CORSConfig{
				AllowOrigins: []string{"*"},
				AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	API.GetAllProducts(e, pool)
	API.GetProductByID(e, pool)
	API.CreateNewProduct(e, pool)
	API.GetAllProductTypes(e, pool)
	API.GetAllColours(e, pool)

	err := e.Start(":8080")

	return err
}
