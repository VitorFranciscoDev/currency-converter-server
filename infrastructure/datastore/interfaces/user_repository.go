package interfaces

import "github.com/VitorFranciscoDev/currency-converter-server/domain/entities"

type UserRepository interface {
	AddUser(user *entities.User) error
}
