package gox_test

import (
	"testing"

	"github.com/goexl/gox"
)

func TestIf(t *testing.T) {
	if 8 != gox.If[int](true, 8, 9) {
		t.Fatalf(`期望：%d，实际：%d`, 8, 9)
	}
}
