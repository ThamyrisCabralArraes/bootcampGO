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
	DeleteDB = "DELETE FROM products WHERE id = ?"
)

type Repository interface {
	Store(product models.Product) (models.Product, error)
	GetOne(id int) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	GetAll() ([]models.Product, error)
	Delete(id int) error
}
type repository struct {
	db *sql.DB
}

func NewRepo(dbConn *sql.DB) Repository {
	return &repository{
		db: dbConn,
	}
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

func (r *repository) GetOne(id int) (models.Product, error) {
	var product models.Product
	db := db.StorageDB

	rows, err := db.Query(GetOneDB, id)
	if err != nil {
		log.Println(err)
		return product, err
	}

	for rows.Next() {
		if err := rows.Scan(&product.ID, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
			log.Println(err)
			return product, err
		}
	}
	return product, nil
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

func (r *repository) Delete(id int) error {
	db := db.StorageDB

	stmt, err := db.Prepare(DeleteDB)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
