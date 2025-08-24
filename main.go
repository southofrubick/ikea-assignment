package main

import (
	"log"
	"github.com/southofrubick/ikea-assignment/db"
)


func main() {
	pool, err := db.InitDB()

	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer pool.Close()
}
