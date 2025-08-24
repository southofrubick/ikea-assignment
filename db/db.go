package db

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

var ctx = context.Background()

func InitDB() (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, "postgres://user:password@localhost:13927/db")

	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Unable to ping database: %w", err)
	}

	err = createTables(pool)
	if err != nil {
		return nil, fmt.Errorf("Unable to create tables: %w", err)
	} else {
		log.Println("All tabler created successfully")
	}

	_, err = populateProductTypesTable(pool)
	if err != nil {
		return nil, fmt.Errorf("Unable to populate product types table: %w", err)
	} else {
		log.Println("Product_Types table populated successfully")
	}

	_, err = populateColoursTable(pool)
	if err != nil {
		return nil, fmt.Errorf("Unable to populate colours table: %w", err)
	} else {
		log.Println("Colours table populated successfully")
	}

	return pool, nil
}

func createTables(db *pgxpool.Pool) error {
	var query = `
	CREATE TABLE IF NOT EXISTS colour (
		id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		name VARCHAR(50) NOT NULL,
		created_at TIMESTAMPTZ DEFAULT NOW(),
		updated_at TIMESTAMPTZ DEFAULT NOW(),
		UNIQUE (name)
	);
	CREATE TABLE IF NOT EXISTS product_type (
		id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		name VARCHAR(50) NOT NULL,
		created_at TIMESTAMPTZ DEFAULT NOW(),
		updated_at TIMESTAMPTZ DEFAULT NOW(),
		UNIQUE (name)
	);
	CREATE TABLE IF NOT EXISTS product (
		id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		name VARCHAR(50) NOT NULL,
		product_type_id INTEGER NOT NULL,
		created_at TIMESTAMPTZ DEFAULT NOW(),
		updated_at TIMESTAMPTZ DEFAULT NOW(),
		CONSTRAINT fk_product_type FOREIGN KEY(product_type_id) REFERENCES product_type(id),
		UNIQUE (name, product_type_id)
	);
	CREATE TABLE IF NOT EXISTS product_colour (
		id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		product_id INTEGER NOT NULL,
		colour_id INTEGER NOT NULL,
		created_at TIMESTAMPTZ DEFAULT NOW(),
		updated_at TIMESTAMPTZ DEFAULT NOW(),
		CONSTRAINT fk_product FOREIGN KEY(product_id) REFERENCES product(id),
		CONSTRAINT fk_colour FOREIGN KEY(colour_id) REFERENCES colour(id),
		UNIQUE (product_id, colour_id)
	);
	`
	_, err := db.Exec(ctx, query)

	if err != nil {
		return fmt.Errorf("error creating table: %w", err)
	}

	return nil
}

func populateProductTypesTable(db *pgxpool.Pool) ([]int, error) {
	file, err := os.Open("product-types.txt")
	if err != nil {
		return nil, fmt.Errorf("error opening product types file: %w", err)
	}
	defer file.Close()

	productTypes, err := io.ReadAll(file)
	productTypesList := strings.Split(string(productTypes[:]), ",")

	var ids []int

	for i := range productTypesList {
		productType := strings.TrimSpace(productTypesList[i])
		var query = `
		INSERT INTO product_type (name)
		VALUES ($1)
		ON CONFLICT (name) DO NOTHING
		RETURNING id;
		`
		id, err := db.Exec(ctx, query, productType)
		if err != nil {
			return nil, fmt.Errorf("error inserting product type: %w", err)
		}
		ids = append(ids, int(id.RowsAffected()))
	}

	return ids, nil
}

func populateColoursTable(db *pgxpool.Pool) ([]int, error) {
	file, err := os.Open("colours.txt")
	if err != nil {
		return nil, fmt.Errorf("error opening colours file: %w", err)
	}
	defer file.Close()

	colours, err := io.ReadAll(file)
	coloursList := strings.Split(string(colours[:]), ",")

	var ids []int

	for i := range coloursList {
		colour := strings.TrimSpace(coloursList[i])
		var query = `
		INSERT INTO colour (name)
		VALUES ($1)
		ON CONFLICT (name) DO NOTHING
		RETURNING id;
		`
		id, err := db.Exec(ctx, query, colour)
		if err != nil {
			return nil, fmt.Errorf("error inserting colour: %w", err)
		}
		ids = append(ids, int(id.RowsAffected()))
	}

	return ids, nil
}
