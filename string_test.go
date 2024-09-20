package gox_test

import (
	"reflect"
	"testing"

	"github.com/goexl/gox"
)

func TestCamel(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"a b", "aB"},
		{"a_b", "aB"},
		{"a-b", "aB"},
		{"Db", "Db"},
	}
	for index, test := range tests {
		camel := gox.String(test.in).Switch().Camel().Build()
		if got := camel.Case(); test.expected != got {
			t.Errorf("第%d个测试未通过，实际：%v，期望：%v", index+1, got, test.expected)
		}
	}
}
func TestStrike(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"a b", "a-b"},
		{"a_b", "a-b"},
		{"a-b", "a-b"},
		{"Db", "db"},
	}
	for index, test := range tests {
		camel := gox.String(test.in).Switch().Strike().Build()
		if got := camel.Case(); test.expected != got {
			t.Errorf("第%d个测试未通过，实际：%v，期望：%v", index+1, got, test.expected)
		}
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		in       string
		expected []string
	}{
		{"a b", []string{"a", "b"}},
		{"你好，世界", []string{"你好，世界"}},
		{"a-b", []string{"a", "b"}},
		{"a_b", []string{"a", "b"}},
		{"helloWorld", []string{"hello", "World"}},
		{"db", []string{"db"}},
	}
	for index, test := range tests {
		split := gox.String(test.in).Split().Naming().Build()
		if got := split.Apply(); !reflect.DeepEqual(test.expected, got) {
			t.Errorf("第%d个测试未通过，实际：%v，期望：%v", index+1, got, test.expected)
		}
	}
}
