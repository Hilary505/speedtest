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
	go speed.ScheduleSpeedTests(results, done, 10*time.Second)
	go speed.WriteResultsToCSV(results, "speedtest_results.csv")
	fmt.Println("Speed test running every 10 seconds. Press Enter to stop...")
	fmt.Scanln()
	// Signal to stop the tests
	done <- true
	time.Sleep(1 * time.Second)
	close(results)
}
