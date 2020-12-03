package gox

import (
	`strings`
)

type StringSlice []string

func (ss *StringSlice) UnmarshalParam(src string) error {
	*ss = strings.Split(src, ",")

	return nil
}
