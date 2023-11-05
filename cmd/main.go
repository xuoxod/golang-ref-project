package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/xuoxod/lab/internal/config"
	"github.com/xuoxod/lab/internal/driver"
	"github.com/xuoxod/lab/internal/envloader"
	"github.com/xuoxod/lab/internal/helpers"
)

// Application configuration
var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := envloader.LoadEnvVars()
	SetupLogs()

	if err != nil {
		infoLog.Println("Error loading environment variables")
		errorLog.Println(err.Error())
	}

	var action string

	flag.StringVar(&action, "act", "", "Sets the command to execute")
	flag.Parse()

	switch action {
	case "testdbc":
		infoLog.Println("Test postgres connection")
		TestDbConn()

	case "environment":
		infoLog.Println("Set development environment")

	case "querydb":
		infoLog.Println("Query postgres datastore")
		fmt.Println(action)

	case "genhash":
		if flag.NArg() < 1 {
			errorLog.Println("Missing argument")
		} else if flag.NArg() > 1 {
			errorLog.Println("Too many arguments")
		} else {
			infoLog.Println("Generate hash string")
			arg := flag.Args()[0]
			hashword, err := helpers.GenerateHash(arg)

			if err != nil {
				errorLog.Println(err.Error())
			}

			fmt.Println(hashword)
		}

	case "comhash":
		infoLog.Println("Compare hash to string")

	case "datesta":
		infoLog.Println("Print date stamp")

	case "timesta":
		infoLog.Println("Print time stamp")

	case "dtstamp":
		infoLog.Println("Print date/time stamp")
	}

	// TestDbConn()

}

func TestDbConn() {
	db, err := run()

	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()
}

func run() (*driver.DB, error) {
	SetupLogs()

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

func SetupLogs() {
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog
}
