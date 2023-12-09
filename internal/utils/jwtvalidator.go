package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/xuoxod/lab/internal/constants"
)

func ValidateJwt(tokenString string) (*jwt.Token, bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SecretKey), nil
	})

	if err != nil {
		return nil, false, err
	}

	return token, token.Valid, nil
}
