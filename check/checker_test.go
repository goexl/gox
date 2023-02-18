package check_test

import (
	"testing"

	"github.com/goexl/gox/check"
)

const (
	checkAll checkType = iota
	checkAny
)

const (
	stringPrefix stringType = iota
	stringSuffix
	stringContains
)

type (
	checkType uint8

	stringType uint8

	stringTest struct {
		t        checkType
		st       stringType
		check    string
		targets  []string
		expected bool
	}
)

func TestString(t *testing.T) {
	tests := []stringTest{
		{t: checkAny, st: stringPrefix, check: "test", targets: []string{"1", "2", "t"}, expected: true},
		{t: checkAll, st: stringPrefix, check: "test", targets: []string{"test", "a", "b"}, expected: false},
		{t: checkAny, st: stringSuffix, check: "1test", targets: []string{"t", "st", "b"}, expected: true},
		{t: checkAll, st: stringSuffix, check: "test", targets: []string{"test1", "a", "b"}, expected: false},
		{t: checkAny, st: stringContains, check: "243test", targets: []string{"2", "24", "b"}, expected: true},
		{t: checkAll, st: stringContains, check: "test", targets: []string{"1test1", "a", "b"}, expected: false},
	}
	for index, test := range tests {
		_check := check.New()
		switch test.t {
		case checkAll:
			_check.All()
		case checkAny:
			_check.Any()
		}

		got := false
		_string := _check.String(test.check).Items(test.targets...)
		switch test.st {
		case stringContains:
			got = _string.Contains().Check()
		case stringPrefix:
			got = _string.Prefix().Check()
		case stringSuffix:
			got = _string.Suffix().Check()
		}

		if test.expected != got {
			t.Errorf("第%d个测试未通过，实际：%v，期望：%v", index+1, got, test.expected)
		}
	}
}
