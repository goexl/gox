package gox

import (
	`testing`
)

func TestRandDigit(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 1, expected: 1},
		{input: 2, expected: 2},
		{input: 3, expected: 3},
		{input: 4, expected: 4},
		{input: 5, expected: 5},
		{input: 6, expected: 6},
	}

	for _, tc := range tests {
		got := RandDigit(tc.input)
		if len(got) != tc.expected {
			t.Fatalf("期望：%v，实际：%v", tc.expected, got)
		}
	}
}

func TestRandString(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 1, expected: 1},
		{input: 2, expected: 2},
		{input: 3, expected: 3},
		{input: 4, expected: 4},
		{input: 5, expected: 5},
		{input: 6, expected: 6},
	}

	for _, tc := range tests {
		got := RandString(tc.input)
		if len(got) != tc.expected {
			t.Fatalf("期望：%v，实际：%v", tc.expected, got)
		}
	}
}
