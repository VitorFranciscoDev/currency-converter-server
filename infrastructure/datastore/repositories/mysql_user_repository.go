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

	var user entities.User

	err := r.db.QueryRow(query, email, password).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *MySQLUserRepository) GetUserByEmail(email string) (*entities.User, error) {
	const query = `
	SELECT * FROM users
	where email = ?;
	`

	var user entities.User

	err := r.db.QueryRow(query, email).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *MySQLUserRepository) AddUser(user *entities.User) (int64, error) {
	const query = `
	INSERT INTO users (uuid, name, email, password) values (?, ?, ?, ?);
	`

	result, err := r.db.Exec(query, user.UUID, user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
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

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
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

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}
