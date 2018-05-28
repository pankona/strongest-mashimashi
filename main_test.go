package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandomChar(t *testing.T) {
	rand.Seed(time.Now().Unix())
	m := make(map[string]int)
	for i := 0; i < 10000; i++ {
		c := randomChar()
		m[c]++
	}
	if len(m) != 26 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(m), 26)
	}
}
