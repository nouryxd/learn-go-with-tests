package main

import "testing"

func TestRacer(t *testing.T) {
	slowURL := "http://facebook.com"
	fastURL := "http://quii.co.uk"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
