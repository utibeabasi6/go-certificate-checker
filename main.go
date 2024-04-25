package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"net/http"
	"time"

	"log/slog"
)

var url string

func init() {
	flag.StringVar(&url, "url", "", "Url to check certifcate expiry")
}

func main() {
	flag.Parse()

	if url == "" {
		slog.Error("No URL specified. Please use the -url flag")
		return
	}

	days, err := fetchExpiryDate(url)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Println("The certificate expires in", *days, "days")
}

func fetchExpiryDate(url string) (*int, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.TLS == nil || len(res.TLS.PeerCertificates) == 0 {
		return nil, errors.New("no certificate found")
	}

	days := int(math.Round(time.Since(res.TLS.PeerCertificates[0].NotAfter).Hours() / 24 * -1))

	return &days, nil
}
