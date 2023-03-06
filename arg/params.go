package arg

type params struct {
	size  int
	short string
	long  string
}

func newParams() *params {
	return &params{
		size:  size,
		short: short,
		long:  long,
	}
}
