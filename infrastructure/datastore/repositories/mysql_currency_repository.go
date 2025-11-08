package repositories

import "database/sql"

type MySQLCurrencyRepository struct {
	db *sql.DB
}

func NewMySQLCurrencyRepository(db *sql.DB) *MySQLCurrencyRepository {
	return &MySQLCurrencyRepository{db}
}
