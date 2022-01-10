package exec

type (
	option interface {
		apply(options *options)
	}

	options struct{}
)

func defaultOptions() *options {
	return &options{}
}
