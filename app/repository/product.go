package repository

import (
	"database/sql"
	"log"
	"nexmedis-technical-test/app/model/entity"
)

type ProductRepository interface {
	GetProducts(name string, limit int, offset int) ([]entity.Product, error)
	GetProductByID(id int) (*entity.Product, error)
	CreateProduct(product entity.Product) error
	UpdateProduct(product entity.Product) error
	DeleteProduct(id int) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) GetProducts(name string, limit int, offset int) ([]entity.Product, error) {
	var rows *sql.Rows
	var err error

	if name != "" {
		// Search by name (case-insensitive)
		query := "SELECT id, name, description, price, stock, category, created_at FROM products WHERE name ILIKE $1 LIMIT $2 OFFSET $3"
		rows, err = r.db.Query(query, "%"+name+"%", limit, offset)
	} else {
		// Fetch all products without filtering by name
		query := "SELECT id, name, description, price, stock, category, created_at FROM products LIMIT $1 OFFSET $2"
		rows, err = r.db.Query(query, limit, offset)
	}

	if err != nil {
		log.Println("Error fetching products:", err)
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Category, &product.CreatedAt); err != nil {
			log.Println("Error scanning product:", err)
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *productRepository) GetProductByID(id int) (*entity.Product, error) {
	var product entity.Product
	err := r.db.QueryRow("SELECT id, name, description, price, stock, category, created_at FROM products WHERE id = $1", id).
		Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock, &product.Category, &product.CreatedAt)
	if err != nil {
		log.Println("Error fetching product by ID:", err)
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) CreateProduct(product entity.Product) error {
	_, err := r.db.Exec("INSERT INTO products (name, description, price, stock, category) VALUES ($1, $2, $3, $4, $5)",
		product.Name, product.Description, product.Price, product.Stock, product.Category)
	if err != nil {
		log.Println("Error inserting product:", err)
	}
	return err
}

func (r *productRepository) UpdateProduct(product entity.Product) error {
	_, err := r.db.Exec("UPDATE products SET name = $1, description = $2, price = $3, stock = $4, category = $5 WHERE id = $6",
		product.Name, product.Description, product.Price, product.Stock, product.Category, product.ID)
	if err != nil {
		log.Println("Error updating product:", err)
	}
	return err
}

func (r *productRepository) DeleteProduct(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		log.Println("Error deleting product:", err)
	}
	return err
}
