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

func (u UserUseCase) Login(email string, password string) (*entities.User, error) {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	return u.r.Login(email, password)
}

func (u UserUseCase) GetUserByEmail(email string) (*entities.User, error) {
	email = strings.TrimSpace(email)

	return u.r.GetUserByEmail(email)
}

func (u UserUseCase) AddUser(user *entities.User) (int64, error) {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	uuid, err := uuid.NewRandom()
	if err != nil {
		return 0, err
	}

	user.UUID = uuid.String()

	return u.r.AddUser(user)
}

func (u UserUseCase) DeleteUser(uuid string) (int64, error) {
	return u.r.DeleteUser(uuid)
}

func (u UserUseCase) UpdateUser(user *entities.User) (int64, error) {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)

	return u.r.UpdateUser(user)
}
