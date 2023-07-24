package main

import (
	"testing"
	"time"
)

func TestSum(t *testing.T) {
	// just a comment
	// new comment
	time.Sleep(time.Second * 3)
	result := sum(1, 3)
	if result != 4 {
		t.Error("got wrong result value!!")
	}
}
