package gox

import (
	"strings"

	"github.com/olivere/elastic/v7"
)

type StringSlice []string

func (ss *StringSlice) UnmarshalParam(src string) error {
	*ss = strings.Split(src, ",")

	return nil
}

func (ss StringSlice) ESSorters() []elastic.Sorter {
	esSorters := make([]elastic.Sorter, 0, len(ss))

	for _, sort := range ss {
		var (
			ascending bool
			tmp       = strings.Split(sort, ":")
		)

		switch len(tmp) {
		case 1:
			if "" == strings.TrimSpace(tmp[0]) {
				continue
			}
		case 2:
			if "ASC" == strings.ToUpper(tmp[1]) {
				ascending = true
			}
		default:
			continue
		}

		esSorters = append(esSorters, elastic.SortInfo{
			Field:     tmp[0],
			Ascending: ascending,
		})
	}

	return esSorters
}
