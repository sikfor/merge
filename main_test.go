package main

import (
	"testing"
	"time"
)

func TestSum(t *testing.T) {
	// just a comment
	time.Sleep(time.Hour * 2)
	result := sum(1, 3)
	if result != 4 {
		t.Error("got wrong result value!!")
	}
}
