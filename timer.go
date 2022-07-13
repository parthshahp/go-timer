package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"gopkg.in/toast.v1"
)

// var notify *notificator.Notificator

func main() {
	// minutes := flag.Int("m", 1, "Number of minutes")
	if len(os.Args[1:]) > 2 {
		os.Exit(1)
	}
	minutes, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Println(err)
	}

	multiple := 1
	if len(os.Args) > 2 {
		multiple, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Println(err)
		}
	}

	for i := 0; i < multiple; i++ {
		runTimer(minutes)
	}

	// fmt.Println("Done")
}

func runTimer(minutes int) {
	seconds := minutes * 60
	for range time.Tick(1 * time.Second) {
		seconds = seconds - 1
		if seconds < 1 {
			break
		}
	}

	notification := toast.Notification{
		AppID:   "Go Timer",
		Title:   "Timer",
		Message: "Total Time: " + strconv.Itoa(minutes) + " minutes",
		Icon:    "D:/go/src/github.com/parthshahp/timer/icon.png", // This file must exist (remove this line if it doesn't)
		Actions: []toast.Action{
			{"protocol", "Ok", ""},
			// {"protocol", "Me too!", ""},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
