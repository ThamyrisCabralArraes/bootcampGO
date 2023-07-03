package util

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb"
	"github.com/google/uuid"
)

func init() {
	txdb.Register("txdb", "mysql", "root@tcp(localhost:3306)/storage")
}

func InitDb() (*sql.DB, error) {
	db, err := sql.Open("txdb", uuid.New().String())

	if err != nil {
		return db, db.Ping()
	}

	return db, err
}
