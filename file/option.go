package file

import (
	`io/fs`
	`os`
)

type (
	option interface {
		apply(options *options)
	}

	options struct {
		_type     _type
		fileMode  fs.FileMode
		writeMode writeMode
		owner     *owner
	}
)

func defaultOptions() *options {
	return &options{
		_type:    TypeFile,
		fileMode: os.ModePerm,
	}
}
