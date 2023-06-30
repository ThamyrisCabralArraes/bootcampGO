package products

import (
	"criando-db/cmd/internal/products/db"
	"criando-db/cmd/internal/products/models"
	"database/sql"
	"log"
)

var (
	StoreDB  = "INSERT INTO products(name, type, count, price) VALUES( ?, ?, ?, ? )"
	GetOneDB = "SELECT * FROM products WHERE id = ?"
	UpdateDB = "UPDATE products SET name = ?, type = ?, count = ?, price = ? WHERE id = ?"
	GetAllDB = "SELECT * FROM products"
)

type Repository interface {
	Store(name, productType string, count int, price float64) (models.Product, error)
	GetOne(id int) models.Product
	Update(product models.Product) (models.Product, error)

	GetAll() ([]models.Product, error)
	Delete(id int) error
}
type repository struct{}

func NewRepo() Repository {
	return &repository{}
}

func (r *repository) Store(product models.Product) (models.Product, error) {
	db := db.StorageDB
	stmt, err := db.Prepare(StoreDB)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var result sql.Result
	result, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price)
	if err != nil {
		return models.Product{}, err
	}
	insertId, _ := result.LastInsertId()
	product.ID = int(insertId)

	return product, nil
}

func (r *repository) GetOne(id int) models.Product {
	var product models.Product
	db := db.StorageDB

	rows, err := db.Query(GetOneDB, id)
	if err != nil {
		log.Println(err)
		return product
	}

	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Println(err)
			return product
		}
	}
	return product
}

func (r *repository) Update(product models.Product) (models.Product, error) {
	db := db.StorageDB

	var products models.Product

	stmt, err := db.Prepare(UpdateDB)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Type, product.Count, product.Price, product.ID)
	if err != nil {
		return products, err
	}
	return product, nil
}

func (r *repository) GetAll() ([]models.Product, error) {
	db := db.StorageDB

	var products []models.Product

	rows, err := db.Query(GetAllDB)
	if err != nil {
		log.Println(err)
		return products, err
	}

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Println(err)
			return products, err
		}
		products = append(products, product)
	}
	return products, nil
}
