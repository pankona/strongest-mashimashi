package phragen

import (
	"net/url"
	"testing"
)

func TestGetNumFromQuery(t *testing.T) {
	tests := []struct {
		inValues  url.Values
		wantNum   int
		wantIsErr bool
	}{
		{nil, 0, true},
		{map[string][]string{"num": {"0"}}, 0, true},
		{map[string][]string{"num": {"-1"}}, 0, true},
		{map[string][]string{"num": {"6"}}, 0, true},
		{map[string][]string{"num": {"-1", "1"}}, 0, true},
		{map[string][]string{"num": {"hoge"}}, 0, true},
		{map[string][]string{"num": {"1"}}, 1, false},
		{map[string][]string{"num": {"5"}}, 5, false},
	}

	for _, tt := range tests {
		n, err := getNumFromQuery(tt.inValues)
		if n != tt.wantNum {
			t.Fatalf("unexpected result. [want] %d [got] %d", n, tt.wantNum)
		}
		if (err != nil) != tt.wantIsErr {
			t.Fatalf("unexpected result. [want] %v [got] %v", err, tt.wantIsErr)
		}
	}
}
