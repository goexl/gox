package gox_test

import (
	"fmt"
	"testing"

	"github.com/goexl/gox"
)

func receiveSlice[T any](data gox.Slice[T]) {
	fmt.Println(data)
}

func receiveIntSlice(data []int) {
	fmt.Println(data)
}

func TestNewSlice(t *testing.T) {
	receiveSlice(gox.NewSlice(1))
	receiveIntSlice(gox.NewSlice(2))
}
