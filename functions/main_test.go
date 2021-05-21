package main

import (
	"net/url"
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

func TestGetNumFromQuery(t *testing.T) {
	tcs := []struct {
		inValues  url.Values
		wantNum   int
		wantIsErr bool
	}{
		{nil,
			0, true},
		{map[string][]string{"num": []string{"0"}},
			0, true},
		{map[string][]string{"num": []string{"-1"}},
			0, true},
		{map[string][]string{"num": []string{"6"}},
			0, true},
		{map[string][]string{"num": []string{"-1", "1"}},
			0, true},
		{map[string][]string{"num": []string{"hoge"}},
			0, true},
		{map[string][]string{"num": []string{"1"}},
			1, false},
		{map[string][]string{"num": []string{"5"}},
			5, false},
	}

	for _, tc := range tcs {
		n, err := getNumFromQuery(tc.inValues)
		if n != tc.wantNum {
			t.Fatalf("unexpected result. [want] %d [got] %d", n, tc.wantNum)
		}
		if (err != nil) != tc.wantIsErr {
			t.Fatalf("unexpected result. [want] %v [got] %v", err, tc.wantIsErr)
		}
	}
}
