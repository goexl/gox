package gox

import (
	"testing"
)

func TestIf(t *testing.T) {
	if 8 != If[int](true, 8, 9) {
		t.Fatalf(`期望：%d，实际：%d`, 8, 9)
	}

	if "true" != If[string](true, "true", "false") {
		t.Fatalf(`期望：%s，实际：%s`, "true", "false")
	}
}
