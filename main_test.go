package main

import (
	"testing"
)

func TestCertificateExpiry(t *testing.T) {
	// Assume googles certificate never expires
	url := "https://google.com"
	expiryDays, err := fetchExpiryDate(url)
	if err != nil {
		t.Fatal(err.Error())
	}

	if *expiryDays <= 0 {
		t.Fatal("Certificate expired")
	}
}
