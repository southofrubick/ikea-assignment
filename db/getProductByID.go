package db

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/southofrubick/ikea-assignment/entity"
)

func GetProductByID(db *pgxpool.Pool, id int) (*entity.Product, error) {
	var query = `
		SELECT id, name, product_type_id, created_at, updated_at FROM product WHERE id = $1;
	`

	var product = &entity.Product{}
	err := db.QueryRow(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.ProductTypeID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error getting product by id: %w", err)
	}

	product.Colours, err = getColoursByProductID(db, product.ID)
	if err != nil {
		return nil, fmt.Errorf("error getting colours for product: %w", err)
	}

	productType, err := getProductTypeByID(db, product.ProductTypeID)
	if err != nil {
		return nil, fmt.Errorf("error getting product_type for product: %w", err)
	}

	product.ProductType = productType

	return product, nil
}

func getColoursByProductID(db *pgxpool.Pool, productId int) ([]string, error) {
	var query = `
		SELECT colour_id FROM product_colour WHERE product_id = $1;
	`

	rows, err := db.Query(ctx, query, productId)
	if err != nil {
		return nil, fmt.Errorf("error getting product_colours for product: %w", err)
	}
	defer rows.Close()

	var colourIds []int

	for rows.Next() {
		var colourId int
		err := rows.Scan(&colourId)

		if err != nil {
			return nil, fmt.Errorf("error scanning product_colours: %w", err)
		}

		colourIds = append(colourIds, colourId)
	}

	query = `
		SELECT name FROM colour WHERE id = ANY ($1);
	`

	rows, err = db.Query(ctx, query, colourIds)
	if err != nil {
		return nil, fmt.Errorf("error getting colours for product: %w", err)
	}
	defer rows.Close()

	var colours []string

	for rows.Next() {
		var colour string
		err := rows.Scan(&colour)

		if err != nil {
			return nil, fmt.Errorf("error scanning colours: %w", err)
		}

		colours = append(colours, colour)
	}

	return colours, nil
}

func getProductTypeByID(db *pgxpool.Pool, id int) (string, error) {
	var query = `
		SELECT name FROM product_type WHERE id = $1;
	`

	var productType string
	err := db.QueryRow(ctx, query, id).Scan(&productType)

	if err != nil {
		return "", fmt.Errorf("error getting product by id: %w", err)
	}

	return productType, nil
}
