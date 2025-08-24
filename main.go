package main

import (
	"log"

	"github.com/southofrubick/ikea-assignment/db"
	"github.com/southofrubick/ikea-assignment/services"
)


func main() {
	pool, err := db.InitDB()

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer pool.Close()

	
	err = services.InitAPI(pool)
	if err != nil {
		log.Fatalf("Failed to start API service: %v", err)
	}
}
