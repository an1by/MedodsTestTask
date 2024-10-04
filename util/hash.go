package util

import (
	"golang.org/x/crypto/bcrypt"
)

func HashToken(token string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(token), 14)
	return string(bytes), err
}

func CompareHash(hashedString string, inputString string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(inputString))
}