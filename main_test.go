package main

import (
	"testing"
)

func TestLoadWords(t *testing.T) {
	noun, err := loadWords("noun.txt", 76216)
	if err != nil {
		t.Fatalf("failed to read words: %s", err.Error())
	}
	if len(noun) != 76216 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(noun), 76216)
	}

	adjective, err := loadWords("adjective.txt", 26664)
	if err != nil {
		t.Fatalf("failed to read words: %s", err.Error())
	}
	if len(adjective) != 26664 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(adjective), 76216)
	}
}
