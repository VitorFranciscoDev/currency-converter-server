package interfaces

import "github.com/VitorFranciscoDev/currency-converter-server/domain/entities"

type UserRepository interface {
	Login(email string, password string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	AddUser(user *entities.User) (int64, error)
	DeleteUser(uuid string) (int64, error)
	UpdateUser(user *entities.User) (int64, error)
}
