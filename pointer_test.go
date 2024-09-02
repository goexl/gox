package gox_test

import (
	"fmt"
	"testing"

	"github.com/goexl/gox"
)

func receive[T any](data gox.Pointer[T]) {
	fmt.Println(data)
}

func receiveInt(data *int) {
	fmt.Println(data)
}

func TestNewPointer(t *testing.T) {
	var testInt = 6
	receive(gox.NewPointer(testInt))
	receiveInt(gox.NewPointer(testInt))
}
