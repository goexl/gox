package tpl_test

import (
	"testing"

	"github.com/goexl/gox/tpl"
)

type (
	user struct {
		Name string
	}

	stringTest struct {
		input    string
		data     any
		expected string
	}
)

func TestString(t *testing.T) {
	user := new(user)
	user.Name = "storezhang"
	tests := []stringTest{
		{input: `名字：{{ .Name }}`, data: user, expected: "名字：storezhang"},
		{input: `欢迎：{{ .Name }}，你好`, data: user, expected: "欢迎：storezhang，你好"},
	}

	for index, test := range tests {
		got, err := tpl.New(test.input).Text().String().Data(test.data).Build().String()
		if nil != err {
			t.Errorf("第%d个测试出错，原因：%v", index+1, err)
		} else if got != test.expected {
			t.Errorf("第%d个测试未通过，实际：%v，期望：%v", index+1, got, test.expected)
		}
	}
}
