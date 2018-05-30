package main

import (
	"testing"
)

func TestReadWords(t *testing.T) {
	noun, err := readWords("noun.txt", 76216)
	if err != nil {
		t.Fatalf("failed to read words: %s", err.Error())
	}
	if len(noun) != 76216 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(noun), 76216)
	}

	adjective, err := readWords("adjective.txt", 26664)
	if err != nil {
		t.Fatalf("failed to read words: %s", err.Error())
	}
	if len(adjective) != 26664 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(adjective), 76216)
	}
}
