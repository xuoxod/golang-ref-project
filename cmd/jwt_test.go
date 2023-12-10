package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/xuoxod/lab/internal/models"
	"github.com/xuoxod/lab/internal/utils"
)

func TestJwtGenerator(t *testing.T) {
	user := models.User{
		FirstName: "rick",
		LastName:  "walker",
		Email:     "ric@email.net",
		Password:  "eatme@jo0es",
		CreatedAt: time.Now(),
	}

	tokenString, err := utils.GenerateJwt(user)

	if err != nil {
		fmt.Println("Error generating the token")
		fmt.Println(err.Error())
		return
	}

	log.Printf("Token:\n\t%s\n", tokenString)
}

func TestJwtValidator(t *testing.T) {
	// time.Sleep(time.Second * 2)

	user := models.User{
		FirstName: "rick",
		LastName:  "walker",
		Email:     "ric@email.net",
		Password:  "eatme@jo0es",
		CreatedAt: time.Now(),
	}

	tokenString, _ := utils.GenerateJwt(user)

	fmt.Printf("\nValidating the token\n")

	token, isValid, err := utils.ValidateToken(tokenString)

	if isValid {
		claims := token.Claims.(*jwt.StandardClaims)
		expiresAt := claims.ExpiresAt
		issuer := claims.Issuer

		fmt.Println("Token:\t", token)
		fmt.Println("Expires At:\t", expiresAt)
		fmt.Println("Issuer:\t", issuer)
		fmt.Printf("\n\n")
	} else {
		fmt.Println(err.Error())
	}
}
