package gox_test

import (
	"strconv"
	"testing"

	"github.com/goexl/gox"
)

func TestFormatInt(t *testing.T) {
	tests := []struct {
		in       int64
		base     uint8
		expected string
	}{
		{in: 1, base: 76, expected: "1"},
		{in: 2, base: 76, expected: "2"},
		{in: 3, base: 76, expected: "3"},
		{in: 4, base: 76, expected: "4"},
		{in: 5, base: 76, expected: "5"},
		{in: 6, base: 76, expected: "6"},
	}

	for _, test := range tests {
		got := gox.FormatInt(test.in, test.base, strconv.FormatInt)
		if got != test.expected {
			t.Fatalf("期望：%v，实际：%v", test.expected, got)
		}
	}
}
func TestFormatUint(t *testing.T) {
	tests := []struct {
		in       uint64
		base     uint8
		expected string
	}{
		{in: 1, base: 76, expected: "1"},
		{in: 2, base: 76, expected: "2"},
		{in: 3, base: 76, expected: "3"},
		{in: 4, base: 76, expected: "4"},
		{in: 5, base: 76, expected: "5"},
		{in: 6, base: 76, expected: "6"},
	}

	for _, test := range tests {
		got := gox.FormatUint(test.in, test.base, strconv.FormatUint)
		if got != test.expected {
			t.Fatalf("期望：%v，实际：%v", test.expected, got)
		}
	}
}

func TestAtoi(t *testing.T) {
	tests := []struct {
		in       string
		base     int
		expected int64
	}{
		{in: "1", base: 76, expected: 1},
		{in: "2", base: 76, expected: 2},
		{in: "3", base: 76, expected: 3},
		{in: "4", base: 76, expected: 4},
		{in: "5", base: 76, expected: 5},
		{in: "6", base: 76, expected: 6},
	}

	for _, test := range tests {
		got := gox.Atoi[int64](test.in, test.base)
		if got != test.expected {
			t.Fatalf("期望：%v，实际：%v", test.expected, got)
		}
	}
}
