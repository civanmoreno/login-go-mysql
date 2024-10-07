package ports

import "main/internal/users/domain"

type UserRepository interface {
	CreateUser(user domain.User) error
	GetUserByUsername(username string) (domain.User, error)
}
