package usecases

import (
	"strings"

	"github.com/VitorFranciscoDev/currency-converter-server/domain/entities"
	"github.com/VitorFranciscoDev/currency-converter-server/infrastructure/datastore/interfaces"
	"github.com/google/uuid"
)

type UserUseCase struct {
	r interfaces.UserRepository
}

func NewUserUseCase(r interfaces.UserRepository) *UserUseCase {
	return &UserUseCase{r: r}
}

func (u UserUseCase) AddUser(user *entities.User) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	user.UUID = uuid.String()

	return u.r.AddUser(user)
}
