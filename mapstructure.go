package gox

import (
	`github.com/mitchellh/mapstructure`
)

type (
	// MapStruct Map和结构体转换
	MapStruct interface {
		// StructToMap 转换成Map
		StructToMap() (model map[string]interface{}, err error)
		// MapToStruct 转换成结构体
		MapToStruct(model map[string]interface{}) (err error)
	}
)

// StructToMap 结构体转换成Map
func StructToMap(obj interface{}, opts ...mapstructOption) (model map[string]interface{}, err error) {
	options := defaultMapstructOptions()
	for _, opt := range opts {
		opt.applyMapstruct(options)
	}

	var decoder *mapstructure.Decoder
	model = make(map[string]interface{})
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

// MapToStruct Map转换成结构体
func MapToStruct(model map[string]interface{}, obj interface{}, opts ...mapstructOption) (err error) {
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
func CopyStruct(from interface{}, to interface{}, opts ...mapstructOption) (err error) {
	tmp := make(map[string]interface{})
	if tmp, err = StructToMap(from, opts...); nil != err {
		return
	}
	err = MapToStruct(tmp, to, opts...)

	return
}
