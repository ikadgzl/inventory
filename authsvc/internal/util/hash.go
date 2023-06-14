package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	return bytes, nil
}

func CheckPasswordHash(hashed []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(hashed, []byte(password))

	return err == nil
}
