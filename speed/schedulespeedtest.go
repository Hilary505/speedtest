package speed

import (
	"fmt"
	"time"
)

type SpeedTestResult struct {
	Timestamp  time.Time
	Download   float64 // in Mbps
	Upload     float64 // in Mbps
	Latency    time.Duration
	ServerName string
}
func ScheduleSpeedTests(results chan<- SpeedTestResult, done <-chan bool, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Run first test immediately
	if err := RunSpeedTest(results); err != nil {
		fmt.Printf("Error running speed test: %v\n", err)
	}

	for {
		select {
		case <-ticker.C:
			if err := RunSpeedTest(results); err != nil {
				fmt.Printf("Error running speed test: %v\n", err)
			}
		case <-done:
			return
		}
	}
}
