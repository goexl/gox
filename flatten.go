package gox

import (
	"encoding/json"
	"strconv"
)

type (
	flattenType interface {
		map[string]any | string
	}

	mapFlatten[T flattenType] struct {
		from   T
		style  *separatorStyle
		prefix string
	}

	separatorStyle struct {
		before string
		middle string
		after  string
	}
)

func Flatten[T flattenType](from T) *mapFlatten[T] {
	return &mapFlatten[T]{
		from:   from,
		style:  &separatorStyle{middle: "."},
		prefix: "",
	}
}

func (mf *mapFlatten[T]) DotStyle() *mapFlatten[T] {
	mf.style = &separatorStyle{middle: "."}

	return mf
}

func (mf *mapFlatten[T]) PathStyle() *mapFlatten[T] {
	mf.style = &separatorStyle{middle: "/"}

	return mf
}

func (mf *mapFlatten[T]) RailsStyle() *mapFlatten[T] {
	mf.style = &separatorStyle{before: "[", after: "]"}

	return mf
}

func (mf *mapFlatten[T]) UnderscoreStyle() *mapFlatten[T] {
	mf.style = &separatorStyle{middle: "_"}

	return mf
}

func (mf *mapFlatten[T]) Prefix(prefix string) *mapFlatten[T] {
	mf.prefix = prefix

	return mf
}

func (mf *mapFlatten[T]) Convert() (to map[string]any, err error) {
	to = make(map[string]any)
	from := any(mf.from)
	switch _from := from.(type) {
	case map[string]any:
		err = mf.flatten(true, &to, _from)
	case string:
		err = mf.string([]byte(_from), &to)
	}

	return
}

func (mf *mapFlatten[T]) string(bytes []byte, to *map[string]any) (err error) {
	from := new(map[string]any)
	if err = json.Unmarshal(bytes, from); nil == err {
		err = mf.flatten(true, to, from)
	}

	return
}

func (mf *mapFlatten[T]) flatten(top bool, flatMap *map[string]any, nested any) (err error) {
	switch nested.(type) {
	case map[string]any:
		for key, value := range nested.(map[string]any) {
			newKey := mf.enKey(top, key)
			err = mf.assign(flatMap, newKey, value)
		}
	case []any:
		for index, value := range nested.([]interface{}) {
			newKey := mf.enKey(top, strconv.Itoa(index))
			err = mf.assign(flatMap, newKey, value)
		}
	}

	return
}

func (mf *mapFlatten[T]) assign(flatMap *map[string]any, newKey string, value any) (err error) {
	switch value.(type) {
	case map[string]any, []any:
		if err := mf.flatten(false, flatMap, value); err != nil {
			return err
		}
	default:
		(*flatMap)[newKey] = value
	}

	return nil
}

func (mf *mapFlatten[T]) enKey(top bool, subKey string) string {
	key := mf.prefix

	if top {
		key += subKey
	} else {
		key += mf.style.before + mf.style.middle + subKey + mf.style.after
	}

	return key
}
