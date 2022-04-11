package gox_test

import (
	"testing"

	"github.com/goexl/gox"
)

func TestParseByte(t *testing.T) {
	tests := []struct {
		in     string
		expect gox.Size
	}{
		{in: `1m`, expect: gox.SizeMB},
		{in: `1g`, expect: gox.SizeGB},
		{in: `1g 1m`, expect: gox.SizeGB + gox.SizeMB},
	}

	for _, test := range tests {
		actual, _ := gox.ParseSize(test.in)
		if actual != test.expect {
			t.Fatalf(`期望：%v，实际：%v`, test.expect, actual)
		}
	}
}

func TestFormatByte(t *testing.T) {
	tests := []struct {
		in     gox.Size
		expect string
	}{
		{in: gox.SizeMB, expect: `1m`},
		{in: gox.SizeGB, expect: `1g`},
		{in: gox.SizeGB + gox.SizeMB, expect: `1g 1m`},
	}

	for _, test := range tests {
		actual := test.in.Format()
		if actual != test.expect {
			t.Fatalf(`期望：%v，实际：%v`, test.expect, actual)
		}
	}
}
