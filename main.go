package main

import (
	"log"
	"net/http"
	"time"
)

const location = "Asia/Tokyo"

func init() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	days, err := sslExpireDays("https://example.com")
	if err != nil {
		log.Println(err)
	}
	log.Printf("days: %d", days)
}

func sslExpireDays(url string) (days int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	expireUTCTime := resp.TLS.PeerCertificates[0].NotAfter
	expireJSTTime := expireUTCTime.In(time.FixedZone("Asia/Tokyo", 9*60*60))

	now := time.Now()

	duration := expireJSTTime.Sub(now)

	hours := int(duration.Hours())
	days = hours / 24
	// hours := hours0 % 24
	return days, err
}
