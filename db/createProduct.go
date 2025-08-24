package db

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func createProductColour(db *pgxpool.Pool, productID int, colourID int) (int, error) {
	var query = `
	INSERT INTO product_colour (product_id, colour_id)
	VALUES ($1, $2)
	RETURNING id;
	`
	var id int
	err := db.QueryRow(ctx, query, productID, colourID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error creating product colour: %w", err)
	}
	return id, nil
}

func CreateProduct(db *pgxpool.Pool, name string, productTypeID int, colourID int) (int, error) {
	var query = `
	INSERT INTO product (name, product_type_id)
	VALUES ($1, $2)
	RETURNING id;
	`

	var id int

	err := db.QueryRow(ctx, query, name, productTypeID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("error creating product: %w", err)
	}

	_, err = createProductColour(db, id, colourID)
	if err != nil {
		return 0, fmt.Errorf("error creating product colour: %w", err)
	}

	return id, nil
}

