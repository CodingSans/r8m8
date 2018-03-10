package sql

import (
	"database/sql"
	"log"
)

// DB interface
type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Begin() (Transaction, error)
}

type db struct {
	db *sql.DB
}

func (d *db) Exec(query string, args ...interface{}) (sql.Result, error) {
	log.Println(query, args)
	return d.db.Exec(query, args...)
}

func (d *db) Query(query string, args ...interface{}) (*sql.Rows, error) {
	log.Println(query, args)
	return d.db.Query(query, args...)
}

func (d *db) QueryRow(query string, args ...interface{}) *sql.Row {
	log.Println(query, args)
	return d.db.QueryRow(query, args...)
}

func (d *db) Begin() (Transaction, error) {
	log.Println("Transaction Begin")
	var transaction Transaction

	tx, err := d.db.Begin()
	if err != nil {
		return transaction, err
	}

	transaction = NewSQLTransaction(tx)
	return transaction, nil
}

// NewSQLDB factory
func NewSQLDB(sqlDB *sql.DB) DB {
	return &db{
		db: sqlDB,
	}
}
