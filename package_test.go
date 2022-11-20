package gox

import (
	"testing"
)

func TestParsePackage(t *testing.T) {
	testcases := []struct {
		input    Package
		expected string
	}{
		{input: "", expected: "com"},
		{input: "a", expected: "a"},
		{input: "www.a.com", expected: "com.a.www"},
		{input: "c.b.a.com", expected: "com.a.b.c"},
		{input: "b-d.a.com", expected: "com.a.d.b"},
		{input: "yunke-web.dev.class100.com", expected: "com.class100.dev.web.yunke"},
	}

	for _, testcase := range testcases {
		got := testcase.input.Name()
		if got != testcase.expected {
			t.Fatalf("期望：%v，实际：%v", testcase.expected, got)
		}
	}
}
