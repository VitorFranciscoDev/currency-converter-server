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

func (r *MySQLUserRepository) Login(email string, password string) (*entities.User, error) {
	const query = `
	SELECT * FROM users
	where email = ? AND password = ?;
	`

	result, err := r.db.Exec(query, email, password)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MySQLUserRepository) GetUserByEmail(email string) (*entities.User, error) {
	const query = `
	SELECT * FROM users
	where email = ?;
	`

	result, err := r.db.Exec(query, email)
	if err != nil {
		return nil, err
	}

	result result, nil
}

func (r *MySQLUserRepository) AddUser(user *entities.User) (int64, error) {
	const query = `
	INSERT INTO users (uuid, name, email, password) values (?, ?, ?, ?);
	`

	result, err := r.db.Exec(query, user.UUID, user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (r *MySQLUserRepository) DeleteUser(uuid string) (int64, error) {
	const query = `
	DELETE FROM users
	WHERE uuid = ?;
	`

	result, err := r.db.Exec(query, uuid)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (r *MySQLUserRepository) UpdateUser(user *entities.User) (int64, error) {
	const query = `
	UPDATE users
	SET
		id = ?,
		uuid = ?,
		name = ?,
		email = ?,
		password = ?,
	WHERE uuid = ?;
	`

	result, err := r.db.Exec(query, user.ID, user.UUID, user.Name, user.Email, user.Password, user.UUID)
	if err != nil {
		return 0, nil
	}

	return result, nil
}
