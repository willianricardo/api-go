package repository

import (
	"api/entity"
	"database/sql"
)

type ProductPostgresRepository struct {
	db *sql.DB
}

func NewProductPostgresRepository(db *sql.DB) *ProductPostgresRepository {
	return &ProductPostgresRepository{
		db: db,
	}
}

func (repository *ProductPostgresRepository) GetProducts() ([]entity.Product, error) {
	products := make([]entity.Product, 0)
	rows, err := repository.db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (repository *ProductPostgresRepository) GetProductByID(id string) (entity.Product, error) {
	rows, err := repository.db.Query("SELECT id, name, price FROM products WHERE id = $1", id)
	if err != nil {
		return entity.Product{}, err
	}

	if rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return entity.Product{}, err
		}
		return product, nil
	}

	return entity.Product{}, nil
}

func (repository *ProductPostgresRepository) CreateProduct(product *entity.Product) error {
	_, err := repository.db.Exec("INSERT INTO products (id, name, price) VALUES ($1, $2, $3)", product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func (repository *ProductPostgresRepository) UpdateProduct(product *entity.Product) error {
	_, err := repository.db.Exec("UPDATE products SET name = $1, price = $2 WHERE id = $3", product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repository *ProductPostgresRepository) DeleteProduct(id string) error {
	_, err := repository.db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
