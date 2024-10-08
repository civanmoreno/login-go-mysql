package utilities

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

// Charset used for generating random passwords.
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"

// GenerateRandomPassword generates a random password of the given length.
func GenerateRandomPassword(length int) (string, error) {
	password := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := range password {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		password[i] = charset[randomIndex.Int64()]
	}
	return string(password), nil
}

// HashPassword hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fmt.Errorf("error al hashear la contraseña: %w", err)
	}

	return string(hashedPassword), nil
}

// ComparePasswords compares the given hashed password with the given password.
func ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
