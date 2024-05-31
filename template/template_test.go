package template_test

import (
	"testing"

	"github.com/goexl/gox/template"
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
	usr := new(user)
	usr.Name = "storezhang"
	tests := []stringTest{
		{input: `名字：{{ .Name }}`, data: usr, expected: "名字：storezhang"},
		{input: `欢迎{{ .Name }}，你好`, data: usr, expected: "欢迎storezhang，你好"},
	}

	for index, test := range tests {
		got, err := template.New(test.input).Text().String().Data(test.data).Build().ToString()
		if nil != err {
			t.Errorf("第%d个测试出错，原因：%v", index+1, err)
		} else if got != test.expected {
			t.Errorf("第%d个测试未通过，实际：%v，期望：%v", index+1, got, test.expected)
		}
	}
}
