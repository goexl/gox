package gox_test

import (
	"fmt"
	"testing"

	"github.com/goexl/gox"
)

func TestStringCheck(t *testing.T) {
	fmt.Println(gox.Checker().Any().String("test").Prefix("testIn", "te").Check())
}
