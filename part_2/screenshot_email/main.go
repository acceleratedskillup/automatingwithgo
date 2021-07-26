package main

import (
	"fmt"
	"os"
	"screenshotemail/email"
	"screenshotemail/screenshot"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
)

const (
	expectedNumOfArgs    = 6
	intervalArgPos       = 1
	emailSenderArgPos    = 2
	passwordArgPos       = 3
	emailRecipientArgPos = 4
	emailSubjectArgPos   = 5
)

/*
interval in minutes
sender email
email password
recipient email
email subject
*/
func main() {
	if len(os.Args) != expectedNumOfArgs {
		fmt.Println("Usage: go run main.go <interval-in-minutes> <email> <password> <recipient> <subject>")
		os.Exit(1)
	}

	interval, err := strconv.Atoi(os.Args[intervalArgPos])
	if err != nil {
		fmt.Printf("Invalid interval: %v\n", err)
		os.Exit(1)
	}

	emailInfo := email.NewEmailInfo(
		os.Args[emailSenderArgPos],
		os.Args[passwordArgPos],
		os.Args[emailRecipientArgPos],
		os.Args[emailSubjectArgPos],
	)

	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(interval).Minutes().WaitForSchedule().Do(func() {
		fmt.Println("cronjob running")
		imageBuffer, err := screenshot.TakeScreenshot()
		if err != nil {
			fmt.Printf("Error occurred taking screenshot: %v\n", err)
		} else {
			email.SendGmail(imageBuffer, emailInfo)
		}
	})
	scheduler.StartBlocking()
}
