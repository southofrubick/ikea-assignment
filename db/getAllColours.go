package db

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/southofrubick/ikea-assignment/entity"
)

func GetAllColours(db *pgxpool.Pool) ([]entity.Colour, error) {
	var query = `
	SELECT id, name, created_at, updated_at
	FROM colour;
	`
	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error getting all product types: %w", err)
	}
	defer rows.Close()

	var colours []entity.Colour
	for rows.Next() {
		var colour entity.Colour
		err := rows.Scan(
			&colour.ID,
			&colour.Name,
			&colour.CreatedAt,
			&colour.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("error scanning product: %w", err)
		}

		colours = append(colours, colour)
	}

	return colours, nil
}
