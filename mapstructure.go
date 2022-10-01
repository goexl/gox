package gox

import (
	"github.com/mitchellh/mapstructure"
)

var _ = StructToForm

type (
	// MapStructure Map和结构体转换
	MapStructure interface {
		// StructToMap 转换成映射
		StructToMap() (model map[string]any, err error)
		// MapToStruct 转换成结构体
		MapToStruct(model map[string]any) (err error)
	}
)

// StructToMap 结构体转换成Map
func StructToMap(obj any, opts ...mapstructOption) (model map[string]any, err error) {
	options := defaultMapstructOptions()
	for _, opt := range opts {
		opt.applyMapstruct(options)
	}

	var decoder *mapstructure.Decoder
	model = make(map[string]any)
	if decoder, err = mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ZeroFields: options.zeroFields,
		Result:     &model,
		Squash:     options.squash,
		TagName:    options.tag,
	}); nil != err {
		return
	}
	err = decoder.Decode(obj)

	return
}

// StructToForm 结构体转换成表单
func StructToForm(obj any, opts ...mapstructOption) (form map[string]string, err error) {
	var model map[string]any
	if model, err = StructToMap(obj, opts...); nil != err {
		return
	}

	// 转换成表单
	form = make(map[string]string)
	for k, v := range model {
		form[k] = v.(string)
	}

	return
}

// MapToStruct Map转换成结构体
func MapToStruct(model map[string]any, obj any, opts ...mapstructOption) (err error) {
	options := defaultMapstructOptions()
	for _, opt := range opts {
		opt.applyMapstruct(options)
	}

	var decoder *mapstructure.Decoder
	if decoder, err = mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ZeroFields: options.zeroFields,
		Result:     obj,
		Squash:     options.squash,
		TagName:    options.tag,
	}); nil != err {
		return
	}
	err = decoder.Decode(model)

	return
}

// CopyStruct 复制结构体数据
func CopyStruct(from any, to any, opts ...mapstructOption) (err error) {
	tmp := make(map[string]any)
	if tmp, err = StructToMap(from, opts...); nil != err {
		return
	}
	err = MapToStruct(tmp, to, opts...)

	return
}
