package db

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"github.com/southofrubick/ikea-assignment/entity"

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
	}

	_, err = populateProductTypesTable(pool)
	if err != nil {
		return nil, fmt.Errorf("Unable to populate product types table: %w", err)
	}

	_, err = populateColoursTable(pool)
	if err != nil {
		return nil, fmt.Errorf("Unable to populate colours table: %w", err)
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

func GetAllProducts(db *pgxpool.Pool) ([]entity.Product, error) {
	var query = `
	SELECT id, name, product_type_id, created_at, updated_at
	FROM product;
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

	product.Colours, err = GetColoursByProductID(db, product.ID)
	if err != nil {
		return nil, fmt.Errorf("error getting colours for product: %w", err)
	}

	productType, err := GetProductTypeByID(db, product.ProductTypeID)
	if err != nil {
		return nil, fmt.Errorf("error getting product_type for product: %w", err)
	}

	product.ProductType = productType

	return product, nil
}

func GetColoursByProductID(db *pgxpool.Pool, productId int) ([]string, error) {
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

func GetProductTypeByID(db *pgxpool.Pool, id int) (string, error) {
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
