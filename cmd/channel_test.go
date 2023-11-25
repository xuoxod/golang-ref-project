package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/xuoxod/lab/internal/collections"
	"github.com/xuoxod/lab/internal/utils"
)

type Message struct {
	User    collections.User
	Message []string
}

func TestSendReceiveChannel(t *testing.T) {
	var msgChan = make(chan Message)
	num, err := utils.GenerateUserDefinedRandomNumber(6, 28)

	if err != nil {
		errorLog.Println(err.Error())
	}

	go Send(num, msgChan)

	_, ok := <-msgChan
	if ok {
		go Receive(msgChan)
	}
	// go utils.ExecuteAfterTime(2, func() { Receive(msgChan) })
	fmt.Scanln()
}

func Send(num int, msg chan Message) {
	fmt.Printf("Random Number:\t%d\n", num)
	log.Printf("Started sending data to %T\n\n", msg)
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
				words = append(words, utils.GenerateWord(12-i))
			}
		}

		message.Message = words
		msg <- message
		time.Sleep(time.Microsecond * 138789)
	}

	defer func() {
		close(msg)
	}()
}

func Receive(msg chan Message) {
	for {
		message, ok := <-msg
		if !ok {
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
	log.Printf("\nFinished receiving data from %T\n", msg)
	utils.ExitProg(0)
}
