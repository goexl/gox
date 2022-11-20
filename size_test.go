package gox_test

import (
	"testing"

	"github.com/goexl/gox"
)

func TestParseByte(t *testing.T) {
	testcases := []struct {
		in     string
		expect gox.Size
	}{
		{in: `1m`, expect: gox.SizeMB},
		{in: `1g`, expect: gox.SizeGB},
		{in: `1g 1m`, expect: gox.SizeGB + gox.SizeMB},
	}

	for _, testcase := range testcases {
		actual, _ := gox.ParseSize(testcase.in)
		if actual != testcase.expect {
			t.Fatalf(`期望：%v，实际：%v`, testcase.expect, actual)
		}
	}
}

func TestFormatByte(t *testing.T) {
	testcases := []struct {
		in     gox.Size
		expect string
	}{
		{in: gox.SizeMB, expect: `1m`},
		{in: gox.SizeGB, expect: `1g`},
		{in: gox.SizeGB + gox.SizeMB, expect: `1g 1m`},
	}

	for _, testcase := range testcases {
		actual := testcase.in.Formatter().Format()
		if actual != testcase.expect {
			t.Fatalf(`期望：%v，实际：%v`, testcase.expect, actual)
		}
	}
}
