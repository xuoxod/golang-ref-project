package main

import (
	_ "embed"
	"log"
	"os"

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
