package utils

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

type function func()

func StringNoSpaces(arg string) bool {
	return !strings.Contains(arg, " ")
}

func Cap(arg string) string {
	var capped string

	for i, c := range strings.Split(arg, "") {
		if i == 0 {
			capped += strings.ToUpper(c)
		} else {
			capped += c
		}
	}

	return capped
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		log.Println("Error:\t", err.Error())
	}
}

func DateTimeStamp() string {
	// dts := fmt.Sprint("Date: ", time.Now())
	// d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)

	d := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()
	hour, minute, second := time.Now().Local().Clock()

	var suffix string
	var strDay string = fmt.Sprintf("%d", day)

	if strings.HasSuffix(strDay, "1") {
		suffix = "st"
	} else if strings.HasSuffix(strDay, "2") {
		suffix = "nd"
	} else if strings.HasSuffix(strDay, "3") {
		suffix = "rd"
	} else {
		suffix = "th"
	}

	return fmt.Sprintf("%v %v%s %v %v:%v:%v", month, day, suffix, year, hour, minute, second)
}

func DateStamp() string {
	d := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()
	var suffix string
	var strDay string = fmt.Sprintf("%d", day)

	if strings.HasSuffix(strDay, "1") {
		suffix = "st"
	} else if strings.HasSuffix(strDay, "2") {
		suffix = "nd"
	} else if strings.HasSuffix(strDay, "3") {
		suffix = "rd"
	} else {
		suffix = "th"
	}

	return fmt.Sprintf("%v %v%s %v", month, day, suffix, year)
}

func TimeStamp() string {
	hour, minute, second := time.Now().Local().Clock()
	return fmt.Sprintf("%v:%v:%v", hour, minute, second)
}

func ExitProg(exitCode int) {
	os.Exit(exitCode)
}

func ComAllArgs() []string {
	return os.Args
}

func CountAllArgs() int {
	return len(ComAllArgs())
}

func ComArgs() []string {
	args := []string{}

	for i, a := range os.Args {
		if i != 0 {
			args = append(args, a)
		}
	}
	return args
}

func CountArgs() int {
	return len(ComArgs())
}

func ExecuteAfterTime(seconds int, f function) {
	duration := time.Duration(seconds) * time.Second
	timer := time.NewTimer(duration)
	<-timer.C
	f()
}

func GenerateRandomNumber(min, max int) (int, error) {
	return min + rand.Intn(max-min), nil
}
