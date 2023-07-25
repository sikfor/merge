package main

import (
	"testing"
	"time"
)

func TestSum(t *testing.T) {
	// just a comment
	// comment three
	time.Sleep(time.Minute * 5)
	time.Sleep(time.Minute * 4)
	result := sum(1, 3)
	if result != 4 {
		t.Error("got wrong result value")
	}
}
