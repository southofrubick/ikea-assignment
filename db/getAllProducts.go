package db

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/southofrubick/ikea-assignment/entity"
)

func GetAllProducts(db *pgxpool.Pool) ([]entity.Product, error) {
	var query = `
	SELECT id, name, product_type_id, created_at, updated_at
	FROM product
	ORDER BY created_at DESC;
	`
	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error getting all products: %w", err)
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.ProductTypeID,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("error scanning product: %w", err)
		}

		products = append(products, product)
	}

	return products, nil
}

