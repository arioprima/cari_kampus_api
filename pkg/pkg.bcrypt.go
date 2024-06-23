package pkg

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashPassword), nil
}

func ComparePassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
