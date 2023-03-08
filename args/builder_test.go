package args_test

import (
	"reflect"
	"testing"

	"github.com/goexl/gox/args"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		args     []any
		expected []string
	}{
		{args: []any{"go", "test"}, expected: []string{"go", "test"}},
		{args: []any{"go", "mod", "tidy"}, expected: []string{"go", "mod", "tidy"}},
	}

	for index, test := range tests {
		got := args.New().Build().Add(test.args...).Build().String()
		if !reflect.DeepEqual(test.expected, got) {
			t.Fatalf("第%d个测试未通过，期望：%v，实际：%v", index+1, test.args, got)
		}
	}
}

func TestLong(t *testing.T) {
	tests := []struct {
		key      string
		value    any
		expected []string
	}{
		{key: "go", value: "mod", expected: []string{"--go=mod"}},
		{key: "go", value: "fmt", expected: []string{"--go=fmt"}},
	}

	for index, test := range tests {
		got := args.New().Build().Arg(test.key, test.value).Build().String()
		if !reflect.DeepEqual(test.expected, got) {
			t.Fatalf("第%d个测试未通过，期望：%v，实际：%v", index+1, test.expected, got)
		}
	}
}
