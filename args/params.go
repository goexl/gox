package args

type params struct {
	capacity int
	short    string
	long     string
	equal    string
}

func newParams() *params {
	return &params{
		capacity: size,
		short:    short,
		long:     long,
		equal:    equal,
	}
}
