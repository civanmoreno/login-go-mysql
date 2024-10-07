package application

import (
	"errors"
	"main/auth"
	"main/internal/users/domain"
	"main/internal/users/ports"
	"main/utilities"
)

type UserService struct {
	userRepository ports.UserRepository
}

func NewUserService(userRepository ports.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) CreateUser(user domain.User) error {
	return u.userRepository.CreateUser(user)
}

func (u *UserService) Login(username, password string) (domain.User, string, error) {
	user, err := u.userRepository.GetUserByUsername(username)
	if err != nil {
		return user, "", err
	}

	if !utilities.ComparePasswords(user.Password, password) {
		return user, "", errors.New("user or password incorrect")
	}
	token, err := auth.GenerateJWT(user.ID, user.Username, user.Role)
	if err != nil {
		return user, "", err
	}
	return user, token, nil
}
