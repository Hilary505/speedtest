package main

import (
	"fmt"
	"time"

	"internetspeed/speed"
)

func main() {
	// Create a channel for speed test results
	results := make(chan speed.SpeedTestResult)
	done := make(chan bool)

	// Start the speed test scheduler in a goroutine
	go speed.ScheduleSpeedTests(results, done, 10*time.Second)

	// Start the CSV writer in a goroutine
	go speed.WriteResultsToCSV(results, "speedtest_results.csv")

	// Wait for user input to exit
	fmt.Println("Speed test running every 10 seconds. Press Enter to stop...")
	fmt.Scanln()

	// Signal to stop the tests
	done <- true
	close(results)
}
