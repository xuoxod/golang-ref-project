package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/xuoxod/lab/internal/constants"
)

func GenerateJwt(user interface{}) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    fmt.Sprintf("%d", user),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(constants.SecretKey))

	if err != nil {
		return "", errors.New("Unabled to generate token")
	}

	return token, nil
}

func GenerateJwtFor(user interface{}, timeFrame int) (string, error) {
	var tf time.Duration = time.Duration(timeFrame)

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    fmt.Sprintf("%d", user),
		ExpiresAt: time.Now().Add(time.Hour * tf).Unix(),
	})

	token, err := claims.SignedString([]byte(constants.SecretKey))

	if err != nil {
		return "", errors.New("Unabled to generate token")
	}

	return token, nil
}
