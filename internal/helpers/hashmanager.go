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

func ComparePassword(hashedText, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(plainText))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("bcrypt mismatched error: ", err)
		return false
	} else if err != nil {
		fmt.Println("bcrypt error: ", err)
		return false
	}

	return true
}
