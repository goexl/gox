package gox

import (
	`testing`
)

func TestParsePackageName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "", expected: "com"},
		{input: "a", expected: "a"},
		{input: "www.a.com", expected: "com.a.www"},
		{input: "c.b.a.com", expected: "com.a.b.c"},
		{input: "b-d.a.com", expected: "com.a.d.b"},
	}

	for _, tc := range tests {
		got := ParsePackageName(tc.input)
		if got != tc.expected {
			t.Fatalf("期望：%v，实际：%v", tc.expected, got)
		}
	}
}
