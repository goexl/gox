package gox_test

import (
	"testing"

	"github.com/goexl/gox"
)

func TestLabels(t *testing.T) {
	tests := []struct {
		in       gox.Labels
		expected string
	}{
		{gox.Labels{"username": "storezhang", "password": "test"}, `{password="test",username="storezhang"}`},
	}
	for index, test := range tests {
		if got := test.in.String(); test.expected != got {
			t.Errorf("第%d个测试未通过，实际：%v，期望：%v", index+1, got, test.expected)
		}
	}
}
