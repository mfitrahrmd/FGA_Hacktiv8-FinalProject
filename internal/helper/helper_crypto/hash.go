package helper_crypto

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func ComparePassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false
	}

	return true
}
