package gox

import (
	`encoding/json`
	`fmt`
)

var _ = NewFieldsException

type (
	FieldsException interface {
		// Fields 返回错误实体
		// 在某些错误下，可能需要返回额外的信息给前端处理
		// 比如，认证错误，需要返回哪些字段有错误
		Fields() Fields
	}

	fieldsException struct {
		// 数据
		fields Fields
	}
)

// NewFieldsException 创建错误
func NewFieldsException(fields ...Field) *fieldsException {
	return &fieldsException{
		fields: fields,
	}
}

func (fe *fieldsException) Fields() Fields {
	return fe.fields
}

func (fe fieldsException) MarshalJSON() (bytes []byte, err error) {
	bytes, err = json.Marshal(fe.fields)

	return
}

func (fe *fieldsException) Error() string {
	return fmt.Sprintf(`data = %s`, fe.fields.String())
}
