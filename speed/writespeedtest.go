package speed

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

func WriteResultsToCSV(results <-chan SpeedTestResult, filename string) {
	// Create or open the CSV file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Printf("Error opening CSV file: %v\n", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header if file is empty
	if stat, err := file.Stat(); err == nil && stat.Size() == 0 {
		header := []string{"Timestamp", "Download (Mbps)", "Upload (Mbps)", "Latency (ms)", "Server"}
		if err := writer.Write(header); err != nil {
			fmt.Printf("Error writing CSV header: %v\n", err)
			return
		}
	}

	// Process incoming results
	for result := range results {
		record := []string{
			result.Timestamp.Format(time.RFC3339),
			strconv.FormatFloat(result.Download, 'f', 2, 64),
			strconv.FormatFloat(result.Upload, 'f', 2, 64),
			result.Latency.String(),
			result.ServerName,
		}

		if err := writer.Write(record); err != nil {
			fmt.Printf("Error writing to CSV: %v\n", err)
			continue
		}
		writer.Flush()

		fmt.Printf("Result recorded: %s - Download: %.2f Mbps, Upload: %.2f Mbps, Latency: %v\n",
			result.Timestamp.Format("15:04:05"),
			result.Download,
			result.Upload,
			result.Latency)
	}
}
