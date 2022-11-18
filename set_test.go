package gox_test

import (
	"testing"

	"github.com/goexl/gox"
)

func TestSet(t *testing.T) {
	_int := 0
	if gox.Set[int](func(value int) { _int = value }, 8, 9); 8 != _int {
		t.Fatalf(`期望：%d，实际：%d`, 8, 9)
	}
}
