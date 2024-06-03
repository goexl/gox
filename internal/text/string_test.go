package text_test

import (
	"reflect"
	"testing"

	"github.com/goexl/gox/internal/text"
)

func TestString(t *testing.T) {
	toStringTests := []struct {
		in       []any
		expected []string
	}{
		{[]any{0, 1, 2}, []string{"0", "1", "2"}},
		{[]any{0.01, 1.01, 2.01}, []string{"0.01", "1.01", "2.01"}},
		{[]any{"test", 1, 2}, []string{"test", "1", "2"}},
		{[]any{"test", 1.01, 2.01}, []string{"test", "1.01", "2.01"}},
		{[]any{"test", 1, 2.01}, []string{"test", "1", "2.01"}},
		{[]any{"test", "a", "b"}, []string{"test", "a", "b"}},
	}

	for _, test := range toStringTests {
		actual := text.ToStrings(test.in...)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("str(%v) = %v；期望：%v", test.in, actual, test.expected)
		}
	}
}
