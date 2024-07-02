package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
)

func StartRecord() {
	erre := godotenv.Load()
	if erre != nil {
		fmt.Println("Cannot Load env")
	}
	rtspURI := os.Getenv("rtspURI")

	var lastRecordedDate string
	var currentCmd *exec.Cmd

	for {
		now := time.Now()

		year := now.Format("2006")
		month := now.Format("01")
		day := now.Format("02")
		currentDate := now.Format("2006-01-02")

		// Check if the date has changed since the last recording
		if currentDate != lastRecordedDate {
			// Stop the current recording if it's running
			if currentCmd != nil && currentCmd.Process != nil {
				err := currentCmd.Process.Kill()
				if err != nil {
					log.Printf("Failed to kill the process: %v", err)
				} else {
					log.Printf("Successfully killed the process for the previous date")
				}
			}

			dir := fmt.Sprintf("recordings/%s/%s", year, month)
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				log.Fatalf("Failed to create directory: %v", err)
			}

			outputFile := fmt.Sprintf("%s/%s.mp4", dir, day)

			if _, err := os.Stat(outputFile); err == nil {
				log.Printf("File %s already exists", outputFile)
				outputFile = fmt.Sprintf("%s/%s_%s.mp4", dir, day, now.Format("15:04:05"))
			}

			cmd := exec.Command("ffmpeg",
				"-i", rtspURI,
				"-c:v", "copy", "-c:a", "aac", "-strict", "experimental",
				outputFile)
			cmd.Stdout = log.Writer()
			cmd.Stderr = log.Writer()

			err = cmd.Start()
			if err != nil {
				log.Printf("FFmpeg command failed to start: %v", err)
			} else {
				currentCmd = cmd
				log.Printf("Started new recording for date: %s", currentDate)
			}

			lastRecordedDate = currentDate
		}

		// Sleep for a minute before checking again
		time.Sleep(1 * time.Minute)
	}
}
