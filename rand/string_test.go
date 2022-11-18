package rand_test

import (
	"testing"

	"github.com/goexl/gox/rand"
)

func TestDigit(t *testing.T) {
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
		got := rand.New().Length(tc.input).Digital().Generate()
		if len(got) != tc.expected {
			t.Fatalf("期望：%v，实际：%v", tc.expected, got)
		}
	}
}

func TestString(t *testing.T) {
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
		got := rand.New().Length(tc.input).Generate()
		if len(got) != tc.expected {
			t.Fatalf("期望：%v，实际：%v", tc.expected, got)
		}
	}
}

func TestCode(t *testing.T) {
	got := rand.New().Code().Generate()
	if 6 != len(got) {
		t.Fatalf("期望：%v，实际：%v", 6, got)
	}
}
