package main

import (
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

var MsgChan = make(chan Message)

type Message struct {
	Name    string
	UUID    string
	Message []string
}

func main() {
	ConfigureApp()
	ChannelTest()
}

func ChannelTest() {
	num, err := utils.GenerateUserDefinedRandomNumber(14, 66)

	if err != nil {
		errorLog.Println(err.Error())
	}

	go Send(num)
	go utils.ExecuteAfterTime(2, func() { Receive() })
	fmt.Scanln()
	close(MsgChan)
}

func Send(num int) {
	fmt.Printf("Random Number:\t%d\n\n", num)

	for i := 1; i < num; i++ {
		var message Message
		message.Name = utils.GenerateName(16)
		message.UUID = utils.GenerateUID()
		words := []string{}

		for j := 1; j <= 11; j++ {
			if j%2 == 0 {
				words = append(words, utils.GenerateWord(i-j))
			} else {
				words = append(words, utils.GenerateWord(j/(2-i)))
			}
		}

		message.Message = words
		MsgChan <- message
		time.Sleep(time.Microsecond * 138789)
	}
}

func Receive() {
	for {
		message := <-MsgChan
		fmt.Printf("%v\n\n", message)
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
