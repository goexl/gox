package file

type (
	nameOption interface {
		applyName(options *nameOptions)
	}

	nameOptions struct {
		_type _type
		ext   string
	}
)

func defaultNameOptions() *nameOptions {
	return &nameOptions{
		_type: TypeFile,
	}
}
