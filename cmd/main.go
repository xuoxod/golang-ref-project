package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/xuoxod/lab/internal/config"
	"github.com/xuoxod/lab/internal/envloader"
	"github.com/xuoxod/lab/internal/models"
	"github.com/xuoxod/lab/internal/utils"
)

// Application configuration
var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	ConfigureApp()
	utils.ClearScreen()
	// utils.Print(utils.GenerateRandomString(17))

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

	time.Sleep(time.Second * 12)

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

func SetupLogs() {
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog
}

func ConfigureApp() {
	err := envloader.LoadEnvVars()

	if err != nil {
		infoLog.Println("Error loading environment variables")
		errorLog.Println(err.Error())
	}

	SetupLogs()
}
