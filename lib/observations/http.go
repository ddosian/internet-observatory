package observations

import (
	"net/http"
	"net/url"
	"time"
)

type HTTPObservation struct {
	URL          *url.URL
	Status       string
	LatencyMS    int64
	TimestampUTC time.Time
}

func ObserveHTTP(url string) (*HTTPObservation, error) {
	startTime := time.Now().UTC()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	latency := time.Since(startTime).Milliseconds()

	if err != nil {
		return nil, err
	}

	observation := &HTTPObservation{
		URL:          resp.Request.URL,
		Status:       resp.Status,
		LatencyMS:    latency,
		TimestampUTC: startTime,
	}

	return observation, nil
}
