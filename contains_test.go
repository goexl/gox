package gox_test

import (
	"testing"

	"github.com/goexl/gox"
)

type contains[T any] struct {
	items    []T
	item     T
	expected bool
}

func TestContains(t *testing.T) {
	tests := []contains[int]{
		{items: []int{1, 2, 3}, item: 1, expected: true},
		{items: []int{1, 2, 3}, item: 8, expected: false},
	}

	for _, test := range tests {
		got := gox.Contains(&test.items, test.item)
		if got != test.expected {
			t.Fatalf("期望：%v，实际：%v", test.expected, got)
		}
	}
}
