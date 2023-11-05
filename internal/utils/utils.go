package utils

import (
	"fmt"
	"log"
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

	return fmt.Sprintf("%v/%v/%v", month, day, year)
}

func DateStamp() string {
	// dts := fmt.Sprint("Date: ", time.Now())

	// d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)

	d := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()

	return fmt.Sprintf("%v/%v/%v", month, day, year)
}

func TimeStamp(strDate string) {
	const layout = "2006-01-02"
	// const layout = "Mon Jan _2 15:04:05 MST 2006"
	// results, err := time.Parse(layout, strDate)
	results, err := time.Parse(layout, strDate)

	if err != nil {
		fmt.Println("TimeStamp error:\t", err.Error())
		return
	}

	fmt.Println("Formatted Date:\t", results)
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
