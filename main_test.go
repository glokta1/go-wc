package main

import "testing"

func TestCountLines(t *testing.T) {
	testCases := []struct {
		filename string
		expected int
	}{
		{"testdata/empty.txt", 0},
		{"testdata/one_line.txt", 1},
		{"testdata/multiple_lines.txt", 3},
	}

	for _, tc := range testCases {
		got := countLines(tc.filename)
		if got != tc.expected {
			t.Errorf("countLines(%q) = %d, want %d", tc.filename, got, tc.expected)
		}
	}
}
func TestCountWords(t *testing.T) {
	testCases := []struct {
		filename string
		expected int
	}{
		{"testdata/empty.txt", 0},
		{"testdata/one_line.txt", 5},
		{"testdata/multiple_lines.txt", 6},
	}

	for _, tc := range testCases {
		got := countWords(tc.filename)
		if got != tc.expected {
			t.Errorf("countWords(%q) = %d, want %d", tc.filename, got, tc.expected)
		}
	}
}

func TestCountCharacters(t *testing.T) {
	testCases := []struct {
		filename string
		expected int
	}{
		{"testdata/empty.txt", 0},
		{"testdata/one_line.txt", 31},
		{"testdata/multiple_lines.txt", 41},
	}

	for _, tc := range testCases {
		got := countCharacters(tc.filename)
		if got != tc.expected {
			t.Errorf("countCharacters(%q) = %d, want %d", tc.filename, got, tc.expected)
		}
	}
}
