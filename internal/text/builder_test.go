package text_test

import (
	"testing"

	"github.com/goexl/gox/internal/text"
)

func TestBuilder(t *testing.T) {
	tests := []struct {
		inputs   []string
		expected string
	}{
		{inputs: []string{"a"}, expected: "a"},
		{inputs: []string{""}, expected: ""},
		{inputs: []string{"a", "b", "c"}, expected: "abc"},
	}

	for _, test := range tests {
		sb := text.NewBuilder()
		for _, input := range test.inputs {
			sb.Append(input)
		}
		got := sb.String()
		if got != test.expected {
			t.Fatalf("期望：%v，实际：%v", test.expected, got)
		}
	}
}
