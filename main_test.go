package main

import (
	"testing"
	"time"
)

// lets add some comment here

func TestSum(t *testing.T) {
	// comment one
	// comment two
	// comment three
	time.Sleep(time.Second * 4)
	result := sum(1, 3)
	if result != 4 {
		t.Error("got wrong result value")
	}
}

// some comment here as well
