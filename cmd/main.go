package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/xuoxod/lab/internal/config"
	"github.com/xuoxod/lab/internal/driver"
	"github.com/xuoxod/lab/internal/envloader"
)

// Application configuration
var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	envloader.LoadEnvVars()
	var environment bool
	flag.BoolVar(&environment, "env", false, "Set development environment")
	flag.Parse()

	app.InProduction = environment

	db, err := run()

	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()

}

func run() (*driver.DB, error) {
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// Connect to database
	log.Println("Connecting to database ...")

	var host string = os.Getenv("DB_HOST")
	var user string = os.Getenv("DB_USER")
	var password string = os.Getenv("DB_PASSWD")
	var dbname string = os.Getenv("DB_NAME")
	var port int = 5432
	var sslmode string = os.Getenv("SSL_MODE")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err := driver.ConnectSql(psqlInfo)

	if err != nil {
		log.Fatal("Cannot connect to database! Dying ...")
	}

	return db, nil
}
