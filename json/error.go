package json

import (
	"fmt"
)

var _ = SyntaxError

type syntaxError struct {
	original []byte
}

func (se syntaxError) Error() string {
	return fmt.Sprintf("错误的JSON语法%q", string(se.original))
}

// SyntaxError JSON语法错误
func SyntaxError(bytes []byte) error {
	return &syntaxError{
		original: bytes,
	}
}
