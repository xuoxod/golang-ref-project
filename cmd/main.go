package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/xuoxod/lab/internal/collections"
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
	User    collections.User
	Message []string
}

func main() {
	ConfigureApp()

	ChannelTest()
}

func ChannelTest() {
	num, err := utils.GenerateUserDefinedRandomNumber(6, 28)

	if err != nil {
		errorLog.Println(err.Error())
	}

	go Send(num)
	go utils.ExecuteAfterTime(2, func() { Receive() })
	fmt.Scanln()
}

func Send(num int) {
	fmt.Printf("Random Number:\t%d\n\n", num)

	for i := 1; i <= num; i++ {
		var message Message
		message.User.FirstName = utils.GenerateName(7)
		message.User.LastName = utils.GenerateName(10)
		message.User.UID = utils.GenerateUID()
		words := []string{}

		for j := 1; j <= 11; j++ {
			if i%2 == 0 {
				message.User.ContactDetails.Email = fmt.Sprintf("%s@email.net", utils.GenerateName(7))
				words = append(words, utils.GenerateWord(10-i))
			} else if i%2 == 1 {
				message.User.ContactDetails.Email = fmt.Sprintf("%s@tagent.mil", utils.GenerateName(11))
				words = append(words, utils.GenerateWord(5))
			} else if i%2 == 2 {
				message.User.ContactDetails.Email = fmt.Sprintf("%s@evil.com", utils.GenerateName(5))
				words = append(words, utils.GenerateWord(7))
			} else {
				message.User.ContactDetails.Email = fmt.Sprintf("%s@oculus.org", utils.GenerateName(4))
				words = append(words, utils.GenerateWord(12-j))
			}
		}

		message.Message = words
		MsgChan <- message
		time.Sleep(time.Microsecond * 138789)
	}

	defer func() {
		close(MsgChan)
	}()
}

func Receive() {
	for {
		message := <-MsgChan
		if len(message.Message) == 0 {
			break
		} else {
			fmt.Println("First Name:\t", message.User.FirstName)
			fmt.Println("Last Name:\t", message.User.LastName)
			fmt.Println("Email:\t        ", message.User.ContactDetails.Email)
			fmt.Println("UID:\t        ", message.User.UID)
			fmt.Println("Message:\t", message.Message)
			fmt.Printf("\n\n")
		}
	}
	utils.ExitProg(0)
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
