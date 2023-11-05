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
	"github.com/xuoxod/lab/internal/utils"
)

// Application configuration
var app config.AppConfig
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	ConfigureApp()

}

func FlagExp() {
	var test1, test2 string

	flag.StringVar(&test1, "test1", "", "Testing flag arguments")
	flag.StringVar(&test2, "test2", "", "Testing flag2 arguments")
	flag.Parse()

	if flag.Parsed() {
		args := utils.ComArgs()
		argsCount := utils.CountArgs()

		switch argsCount {
		case 1:
			fmt.Println(args)

		case 2:
			fmt.Println(args)

		default:
		}
	}
}

func RunProg() {
	var genhash, comhash, envir, querydb, testdbc, datesta, timesta, dtstamp string

	flag.StringVar(&genhash, "genhash", "", "Generates a hash from given string")
	flag.StringVar(&comhash, "comhash", "", "Compare hash to string")
	flag.StringVar(&envir, "envir", "", "Set development environment")
	flag.StringVar(&querydb, "querydb", "", "Query datastore")
	flag.StringVar(&testdbc, "testdbc", "", "Test database connection")
	flag.StringVar(&datesta, "datesta", "", "Print date stamp")
	flag.StringVar(&timesta, "timesta", "", "Print time stamp")
	flag.StringVar(&dtstamp, "dtstamp", "", "Print date/time stamp")
	flag.Parse()

	flag.Visit(func(f *flag.Flag) {
		name := f.Name
		value := f.Value.String()
		numArgs := flag.NArg()

		switch name {
		case "genhash":

			if flag.Parsed() {
				fmt.Println("Value: ", f.Value.String())
				fmt.Println("Name: ", f.Name)
				fmt.Println("Count: ", flag.NArg())
				fmt.Println("Count Args: ", utils.CountArgs())
			}

			if value != "" && numArgs == 0 {
				infoLog.Println("Generate hash string")
				value := fmt.Sprintf("%v", value)
				hash, err := helpers.GenerateHash(value)

				if err != nil {
					errorLog.Println(err.Error())
				}

				fmt.Println("Original: ", value)
				fmt.Println(hash)
			} else if value != "" && numArgs > 0 {
				errorLog.Println("Too many arguments")
			} else {
				errorLog.Println("Missing argument")
			}
		case "comhash":
			fmt.Println("Value: ", f.Value.String())
			fmt.Println("Name: ", f.Name)
			fmt.Println("Count: ", flag.NArg())

			infoLog.Println("Compare hash to string")
		case "envir":
			infoLog.Println("Set development environment")
		case "querydb":
			infoLog.Println("Query datastore")
		case "testdbc":
			infoLog.Println("Test database connection")
		case "datesta":
			infoLog.Println("Print date stamp")
		case "timesta":
			infoLog.Println("Print time stamp")
		case "dtstamp":
			infoLog.Println("Print date/time stamp")
		}
	})
}

func TestDbConn() {
	db, dsn, err := connectDatastore()

	if err != nil {
		log.Fatal(err)
	}

	app.DBConnection = dsn

	defer db.SQL.Close()
}

func connectDatastore() (*driver.DB, string, error) {
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
		return nil, "", err
	}

	return db, psqlInfo, nil
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
