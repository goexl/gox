package gox

import (
	`testing`
)

func TestFilename(t *testing.T) {
	tests := []struct {
		filename string
		expected string
	}{
		{filename: "/usr/opt/test.mp4", expected: "test"},
		{filename: "/usr/opt/test", expected: "test"},
		{filename: "C:\\Program Files\\test.mp4", expected: "test"},
		{filename: "C:\\Program Files\\test", expected: "test"},
	}

	for _, tc := range tests {
		got := Filename(tc.filename)
		if got != tc.expected {
			t.Fatalf("期望：%v，实际：%v", tc.expected, got)
		}
	}
}
