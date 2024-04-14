package main

import "testing"

func TestCountLines(t *testing.T) {
	testCases := []struct {
		filename string
		expected int
	}{
		{"testdata/empty.txt", 0},
		{"testdata/one_line.txt", 1},
		{"testdata/multiple_lines.txt", 4},
	}

	for _, tc := range testCases {
		got := countLines(tc.filename)
		if got != tc.expected {
			t.Errorf("countLines(%q) = %d, want %d", tc.filename, got, tc.expected)
		}
	}
}
