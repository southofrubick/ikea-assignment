package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/southofrubick/ikea-assignment/db"
)


func main() {
	pool, err := db.InitDB()

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer pool.Close()

	e := echo.New()

	e.Logger.Fatal(e.Start(":8080"))
}
