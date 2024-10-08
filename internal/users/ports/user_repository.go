package ports

import "main/internal/users/domain"

type UserRepository interface {
	CreateUser(user domain.User) error
	GetUserByUsername(username string) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	UpdatePassword(email, password string) error
}
