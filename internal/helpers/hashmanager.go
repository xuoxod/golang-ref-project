package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(characters string) (string, error) {
	hashword, err := bcrypt.GenerateFromPassword([]byte(characters), 12)

	if err != nil {
		return "", err
	}

	return string(hashword), nil
}

func ComparePassword(testPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println(err)
		return false
	} else if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
