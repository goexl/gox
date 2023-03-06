package arg_test

import (
	"reflect"
	"testing"

	"github.com/goexl/gox/arg"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		args     []any
		expected arg.Args
	}{
		{args: []any{"go", "test"}, expected: arg.Args{"go", "test"}},
		{args: []any{"go", "mod", "tidy"}, expected: arg.Args{"go", "mod", "tidy"}},
	}

	for index, test := range tests {
		got := arg.New().Build().Add(test.args...).Build()
		if !reflect.DeepEqual(test.expected, got) {
			t.Fatalf("第%d个测试未通过，期望：%v，实际：%v", index+1, test.args, got)
		}
	}
}

func TestLong(t *testing.T) {
	tests := []struct {
		key      any
		value    any
		expected arg.Args
	}{
		{key: "go", value: "mod", expected: arg.Args{"--go=mod"}},
		{key: "go", value: "fmt", expected: arg.Args{"--go=fmt"}},
	}

	for index, test := range tests {
		got := arg.New().Build().Long(test.key, test.value).Build()
		if !reflect.DeepEqual(test.expected, got) {
			t.Fatalf("第%d个测试未通过，期望：%v，实际：%v", index+1, test.expected, got)
		}
	}
}
