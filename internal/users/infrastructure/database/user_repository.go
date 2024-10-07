package database

import (
	"database/sql"
	"errors"
	"main/internal/users/domain"
	"main/internal/users/ports"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) ports.UserRepository {
	return &MySQLUserRepository{db: db}
}
func (r *MySQLUserRepository) CreateUser(user domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?)", user.Username, user.Email, user.Password, user.Role)
	return err
}

func (r *MySQLUserRepository) GetUserByUsername(username string) (domain.User, error) {
	var user domain.User

	err := r.db.QueryRow("SELECT id, username, email, password, role FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user, errors.New("user or password incorrect")
		}
		return user, err
	}

	return user, nil
}
