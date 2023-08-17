package main

import (
	"flag"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	argument := flag.String("argument", "", "The first argument")
	action := flag.String("action", "", "The action to perform with the first argument")

	flag.Parse()

	switch *action {
	case "hash":
		hashed, err := HashPassword(*argument)

		if err != nil {
			fmt.Printf("\n\tError for action [%s]\n\tUnable to hash the argument [%s]\n\n", *action, *argument)
			return
		}

		fmt.Printf("Original:\t%s\nHashed:\t%s\n\n", *argument, hashed)

	case "compare":
		hashed, err := HashPassword(*argument)

		if err != nil {
			fmt.Printf("\n\tError for action [%s]\n\tUnable to hash the argument [%s]\n\n", *action, *argument)
			return
		}

		results := ComparePassword(*argument, hashed)

		fmt.Printf("Original:\t%s\nHashed:\t%s\n[%s] === [%v]? %v\n\n", *argument, hashed, *argument, hashed, results)

	default:
		fmt.Printf("\nUnknown action argument: [%s]\n\tProgram Ended\n\n", *action)
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePassword(testPassword, hashedPassword string) bool {
	byteHashedPassword := []byte(hashedPassword)
	bytePassword := []byte(testPassword)

	err := bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	if err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println(err)
		return false
	} else if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
