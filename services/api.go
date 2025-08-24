package services

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/southofrubick/ikea-assignment/services/api"
)

func InitAPI(pool *pgxpool.Pool) error {
	e := echo.New()

	api.GetAllProducts(e, pool)
	api.GetProductByID(e, pool)
	api.CreateNewProduct(e, pool)

	err := e.Start(":8080")

	return err
}
