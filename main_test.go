package main

import (
	"testing"
	"time"
)

func TestSum(t *testing.T) {
	// just a comment
	// add a new comment
	time.Sleep(time.Minute * 2)
	result := sum(1, 3)
	if result != 4 {
		t.Error("got wrong result value!!")
	}
}
