package gox_test

import (
	"reflect"
	"testing"

	"github.com/goexl/gox"
)

const (
	dotStyle style = iota
	pathStyle
	railsStyle
	underscoreStyle
)

type style int

func TestFlatten(test *testing.T) {
	testcases := []struct {
		in     map[string]any
		want   map[string]any
		prefix string
		style  style
	}{
		{
			map[string]any{
				"foo": map[string]any{
					"jim": "bean",
				},
				"fee": "bar",
				"n1": map[string]any{
					"alist": []any{
						"a",
						"b",
						"c",
						map[string]any{
							"d": "other",
							"e": "another",
						},
					},
				},
				"number": 1.4567,
				"bool":   true,
			},
			map[string]any{
				"foo.jim":      "bean",
				"fee":          "bar",
				"n1.alist.0":   "a",
				"n1.alist.1":   "b",
				"n1.alist.2":   "c",
				"n1.alist.3.d": "other",
				"n1.alist.3.e": "another",
				"number":       1.4567,
				"bool":         true,
			},
			"",
			dotStyle,
		},
		{
			map[string]any{
				"foo": map[string]any{
					"jim": "bean",
				},
				"fee": "bar",
				"n1": map[string]any{
					"alist": []any{
						"a",
						"b",
						"c",
						map[string]any{
							"d": "other",
							"e": "another",
						},
					},
				},
			},
			map[string]any{
				"foo[jim]":        "bean",
				"fee":             "bar",
				"n1[alist][0]":    "a",
				"n1[alist][1]":    "b",
				"n1[alist][2]":    "c",
				"n1[alist][3][d]": "other",
				"n1[alist][3][e]": "another",
			},
			"",
			railsStyle,
		},
		{
			map[string]any{
				"foo": map[string]any{
					"jim": "bean",
				},
				"fee": "bar",
				"n1": map[string]any{
					"alist": []any{
						"a",
						"b",
						"c",
						map[string]any{
							"d": "other",
							"e": "another",
						},
					},
				},
				"number": 1.4567,
				"bool":   true,
			},
			map[string]any{
				"foo/jim":      "bean",
				"fee":          "bar",
				"n1/alist/0":   "a",
				"n1/alist/1":   "b",
				"n1/alist/2":   "c",
				"n1/alist/3/d": "other",
				"n1/alist/3/e": "another",
				"number":       1.4567,
				"bool":         true,
			},
			"",
			pathStyle,
		},
		{
			map[string]any{
				"a": map[string]any{
					"b": "c",
				},
				"e": "f",
			},
			map[string]any{
				"p:a.b": "c",
				"p:e":   "f",
			},
			"p:",
			dotStyle,
		},
		{
			map[string]any{
				"foo": map[string]any{
					"jim": "bean",
				},
				"fee": "bar",
				"n1": map[string]any{
					"alist": []any{
						"a",
						"b",
						"c",
						map[string]any{
							"d": "other",
							"e": "another",
						},
					},
				},
				"number": 1.4567,
				"bool":   true,
			},
			map[string]any{
				"foo_jim":      "bean",
				"fee":          "bar",
				"n1_alist_0":   "a",
				"n1_alist_1":   "b",
				"n1_alist_2":   "c",
				"n1_alist_3_d": "other",
				"n1_alist_3_e": "another",
				"number":       1.4567,
				"bool":         true,
			},
			"",
			underscoreStyle,
		},
	}

	for index, testcase := range testcases {
		converter := gox.Flatten(testcase.in).Prefix(testcase.prefix)
		switch testcase.style {
		case dotStyle:
			converter.DotStyle()
		case railsStyle:
			converter.RailsStyle()
		case pathStyle:
			converter.PathStyle()
		case underscoreStyle:
			converter.UnderscoreStyle()
		}
		if got, err := converter.Convert(); nil != err {
			test.Errorf("第%d个测试出错%v", index+1, err)
		} else if !reflect.DeepEqual(got, testcase.want) {
			test.Errorf("第%d个测试未通过，实际：%v，期望：%v", index+1, got, testcase.want)
		}
	}
}
