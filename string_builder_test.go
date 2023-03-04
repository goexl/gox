package gox_test

import (
	"testing"

	"github.com/goexl/gox"
)

func TestStringBuilder(t *testing.T) {
	tests := []struct {
		inputs   []string
		expected string
	}{
		{inputs: []string{"a"}, expected: "a"},
		{inputs: []string{""}, expected: ""},
		{inputs: []string{"a", "b", "c"}, expected: "abc"},
	}

	for _, test := range tests {
		sb := gox.StringBuilder()
		for _, input := range test.inputs {
			sb.Append(input)
		}
		got := sb.String()
		if got != test.expected {
			t.Fatalf("期望：%v，实际：%v", test.expected, got)
		}
	}
}
