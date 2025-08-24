package db

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/southofrubick/ikea-assignment/entity"
)

func GetAllProductTypes(db *pgxpool.Pool) ([]entity.ProductType, error) {
	var query = `
	SELECT id, name, created_at, updated_at
	FROM product_type;
	`
	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error getting all product types: %w", err)
	}
	defer rows.Close()

	var productsTypes []entity.ProductType
	for rows.Next() {
		var productType entity.ProductType
		err := rows.Scan(
			&productType.ID,
			&productType.Name,
			&productType.CreatedAt,
			&productType.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("error scanning product: %w", err)
		}

		productsTypes = append(productsTypes, productType)
	}

	return productsTypes, nil
}

