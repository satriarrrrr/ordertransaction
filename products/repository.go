package products

import (
	"database/sql"
	"fmt"
)

var (
	tableProducts = "products"
)

// IRepository an interface
type IRepository interface {
	GetProducts(limit uint16) ([]Product, error)
	GetProductByID(id uint64) (Product, error)
	InsertProduct(product Product) (Product, error)
}

// Repository implements IRepository
type Repository struct {
	DB *sql.DB
}

// NewProductsRepository create repository
func NewProductsRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

// GetProducts get products
func (r *Repository) GetProducts(limit uint16) ([]Product, error) {
	var products []Product
	rows, err := r.DB.Query(
		fmt.Sprintf(
			"SELECT id, name, quantity, price FROM %s WHERE deleted_at IS NULL LIMIT ?",
			tableProducts,
		),
		limit,
	)
	if err != nil {
		return products, err
	}
	defer rows.Close()
	// Appending result
	for rows.Next() {
		var product Product
		if err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity); err != nil {
			return products, err
		}
		products = append(products, product)
	}
	return products, err
}

// GetProductByID get products
func (r *Repository) GetProductByID(id uint64) (Product, error) {
	product := Product{ID: id}
	err := r.DB.QueryRow(
		fmt.Sprintf(
			"SELECT name, quantity, price FROM %s WHERE id = ? AND deleted_at IS NULL",
			tableProducts,
		),
		id,
	).Scan(&product.Name, &product.Quantity, &product.Price)
	return product, err
}

// InsertProduct get products
func (r *Repository) InsertProduct(product Product) (Product, error) {
	stmt, err := r.DB.Prepare(fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		tableProducts,
		"name, quantity, price",
		"?, ?, ?",
	))
	if err != nil {
		return product, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(product.Name, product.Quantity, product.Price)
	if err != nil {
		return product, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return product, err
	}
	product.ID = uint64(id)
	return product, nil
}
