package speed

import (
	"fmt"
	"time"

	"github.com/showwin/speedtest-go/speedtest"
)

func RunSpeedTest(results chan<- SpeedTestResult) error {
	// Initialize speed test client
	client := speedtest.New()

	// Fetch server list
	serverList, err := client.FetchServers()
	if err != nil {
		return fmt.Errorf("failed to fetch servers: %w", err)
	}

	// Find the best server
	targets, err := serverList.FindServer([]int{})
	if err != nil {
		return fmt.Errorf("failed to find server: %w", err)
	}
	bestServer := targets[0]

	// Run tests
	fmt.Println("Running speed test...")
	err = bestServer.DownloadTest()
	if err != nil {
		return fmt.Errorf("download test failed: %w", err)
	}

	err = bestServer.UploadTest()
	if err != nil {
		return fmt.Errorf("upload test failed: %w", err)
	}

	// Ping test (latency)
	var latency time.Duration
	err = bestServer.PingTest(func(l time.Duration) {
		latency = l
	})
	if err != nil {
		return fmt.Errorf("ping test failed: %w", err)
	}

	// Convert ByteRate to Mbps (1 Byte = 8 bits, 1 Mbps = 1,000,000 bits)
	downloadMbps := float64(bestServer.DLSpeed) * 8 / 1000000
	uploadMbps := float64(bestServer.ULSpeed) * 8 / 1000000

	// Send result to channel
	results <- SpeedTestResult{
		Timestamp:  time.Now(),
		Download:   downloadMbps,
		Upload:     uploadMbps,
		Latency:    latency,
		ServerName: bestServer.Name,
	}

	return nil
}