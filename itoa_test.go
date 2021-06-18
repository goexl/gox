package gox

import (
	`testing`
)

func TestFormatInt(t *testing.T) {
	tests := []struct {
		original int64
		base     int
		expected string
	}{
		{original: 1, base: 76, expected: "1"},
		{original: 2, base: 76, expected: "2"},
		{original: 3, base: 76, expected: "3"},
		{original: 4, base: 76, expected: "4"},
		{original: 5, base: 76, expected: "5"},
		{original: 6, base: 76, expected: "6"},
	}

	for _, tc := range tests {
		got := FormatInt(tc.original, tc.base)
		if got != tc.expected {
			t.Fatalf("期望：%v，实际：%v", tc.expected, got)
		}
	}
}

func TestAtoi(t *testing.T) {
	tests := []struct {
		original string
		base     int
		expected int64
	}{
		{original: "1", base: 76, expected: 1},
		{original: "2", base: 76, expected: 2},
		{original: "3", base: 76, expected: 3},
		{original: "4", base: 76, expected: 4},
		{original: "5", base: 76, expected: 5},
		{original: "6", base: 76, expected: 6},
	}

	for _, tc := range tests {
		got := Atoi(tc.original, tc.base)
		if got != tc.expected {
			t.Fatalf("期望：%v，实际：%v", tc.expected, got)
		}
	}
}
