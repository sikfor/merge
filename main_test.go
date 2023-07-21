package main

import (
	"testing"
	"time"
)

func TestSum(t *testing.T) {
	// just a comment
	time.Sleep(time.Minute * 3)
	result := sum(1, 3)
	if result != 5 {
		t.Error("got wrong result value")
	}
}
