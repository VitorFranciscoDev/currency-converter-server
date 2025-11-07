package repositories

import (
	"database/sql"

	"github.com/VitorFranciscoDev/currency-converter-server/domain/entities"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db}
}

func (r *MySQLUserRepository) AddUser(user entities.User) error {
	const query = `
	INSERT INTO users (uuid, name, email, password) values (?, ?, ?, ?)
	`

	_, err := r.db.Exec(query, user.UUID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
