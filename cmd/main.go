package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/xuoxod/lab/internal/config"
	"github.com/xuoxod/lab/internal/envloader"
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
	currentTime := time.Now().Add(time.Hour * 24).Unix()
	// strCurrentTime := fmt.Sprintf("%dµs", currentTime)
	durationHours := time.Duration.Hours(time.Duration(currentTime))

	fmt.Println("Current Time:\t", currentTime)
	fmt.Printf("Type:\t%T\n", currentTime)
	fmt.Println("Duration Hours:\t", durationHours)

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
