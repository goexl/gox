package gox

import (
	`github.com/mitchellh/mapstructure`
)

type (
	// MapStruct Map和结构体转换
	MapStruct interface {
		// ToMap 转换成Map
		StructToMap() (model map[string]interface{}, err error)
		// 转换成结构体
		MapToStruct(model map[string]interface{}) (err error)
	}
)

// StructToMap 结构体转换成Map
func StructToMap(obj interface{}) (model map[string]interface{}, err error) {
	var decoder *mapstructure.Decoder

	model = make(map[string]interface{})
	if decoder, err = mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ZeroFields: false,
		Result:     &model,
		TagName:    "json",
	}); nil != err {
		return
	}

	err = decoder.Decode(obj)

	return
}

// MapToStruct Map转换成结构体
func MapToStruct(model map[string]interface{}, obj interface{}) (err error) {
	var decoder *mapstructure.Decoder

	if decoder, err = mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ZeroFields: false,
		Result:     obj,
		TagName:    "json",
	}); nil != err {
		return
	}

	err = decoder.Decode(model)

	return
}
