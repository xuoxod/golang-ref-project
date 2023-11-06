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
	RunProg()
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
	var genhash, comhash, envir, querydb, testdbc string

	flag.StringVar(&genhash, "genhash", "", "Generates a hash from given string")
	flag.StringVar(&comhash, "comhash", "", "Compare plain text to hashed text. Must embed the hashed text within single quotes.")
	flag.StringVar(&envir, "envir", "", "Set development environment")
	flag.StringVar(&querydb, "querydb", "", "Query datastore")
	flag.StringVar(&testdbc, "testdbc", "", "Test database connection")
	flag.Bool("datesta", true, "Print date stamp")
	flag.Bool("timesta", true, "Print time stamp")
	flag.Bool("dtstamp", true, "Print date/time stamp")
	flag.Parse()

	flag.Visit(func(f *flag.Flag) {
		command := f.Name
		arguments := utils.ComArgs()
		numArgs := len(arguments)

		switch command {
		case "genhash":
			if arguments[1] != "" && numArgs == 2 {
				infoLog.Println("Generate hash string")

				if flag.Parsed() {
					fmt.Println("Command: ", command)
					fmt.Println("Argument: ", arguments[1])
					fmt.Println("Cmd line args: ", utils.CountArgs())
				}

				hash, err := helpers.GenerateHash(arguments[1])

				if err != nil {
					errorLog.Println(err.Error())
				}

				fmt.Println("Original: ", arguments[1])
				fmt.Println("Hashed:   ", hash)
			}
		case "comhash":
			if arguments[1] != "" && arguments[2] != "" && numArgs == 3 {
				infoLog.Println("Compare hash to string")
				arg1 := arguments[1]
				arg2 := arguments[2]

				fmt.Println("Argument 1: ", arg1)
				fmt.Println("Argument 2: ", arg2)

				results := helpers.ComparePassword(arguments[2], arguments[1])

				fmt.Printf("%s === %s? %t\n", arguments[1], arguments[2], results)
			}
		case "envir":
			infoLog.Println("Set development environment")
		case "querydb":
			infoLog.Println("Query datastore")
		case "testdbc":
			infoLog.Println("Test database connection")
		case "datesta":
			if numArgs == 1 {
				infoLog.Println("Print date stamp")
				fmt.Println(utils.DateStamp())
			}
		case "timesta":
			infoLog.Println("Print time stamp")
			fmt.Println(utils.TimeStamp())
		case "dtstamp":
			infoLog.Println("Print date/time stamp")
			fmt.Println(utils.DateTimeStamp())
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
