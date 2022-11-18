package gox

import (
	"encoding/json"
	"fmt"
	"strings"
)

// 本来应该使用exc库，但是会导致包循环依赖，所有复制一份代码内部使用
type fieldError struct {
	message string
	field   Field[any]
}

func newField(message string, field Field[any]) *fieldError {
	return &fieldError{
		message: message,
		field:   field,
	}
}

func (f *fieldError) Message() string {
	return f.message
}

func (f *fieldError) Field() Field[any] {
	return f.field
}

func (f *fieldError) MarshalJSON() (bytes []byte, err error) {
	output := make(map[string]interface{})
	output[`message`] = f.message
	output[f.Field().Key()] = f.Field().Value()
	bytes, err = json.Marshal(output)

	return
}

func (f *fieldError) Error() string {
	var sb strings.Builder
	sb.WriteRune('{')
	sb.WriteString(fmt.Sprintf(`message = %s, `, f.message))
	sb.WriteString(fmt.Sprintf(`%s = %v`, f.field.Key(), f.field.Value()))
	sb.WriteRune('}')

	return sb.String()
}
