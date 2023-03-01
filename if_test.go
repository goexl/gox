package gox_test

import (
	"testing"

	"github.com/goexl/gox"
)

func TestIf(t *testing.T) {
	ift := gox.If(true, 8)
	if 8 != ift {
		t.Fatalf("期望：%d，实际：%d", 8, ift)
	}

	sft := gox.If(true, "true")
	if "true" != sft {
		t.Fatalf(`期望：%s，实际：%s`, "true", sft)
	}
}

func TestIfx(t *testing.T) {
	if 8 != gox.Ternary(true, 8, 9) {
		t.Fatalf(`期望：%d，实际：%d`, 8, 9)
	}

	if "true" != gox.Ternary(true, "true", "false") {
		t.Fatalf(`期望：%s，实际：%s`, "true", "false")
	}
}
