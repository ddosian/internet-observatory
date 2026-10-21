package main

import (
	"fmt"
	"internet-observatory/lib/observations"
	"time"
)

func main() {
	observation, err := observations.ObserveHTTP("https://google.com/")

	if err != nil {
		fmt.Println("Error whilst performing observation:", err)
	}

	fmt.Printf("URL: %s\n", observation.URL.String())
	fmt.Printf("Status: %s\n", observation.Status)
	fmt.Printf("Latency: %d ms\n", observation.LatencyMS)
	fmt.Printf("Timestamp (UTC): %s\n",
		observation.TimestampUTC.Format(time.RFC3339Nano))
}
