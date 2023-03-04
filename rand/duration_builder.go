package rand

type durationBuilder struct {
	params *durationParams
}

func newDurationBuilder() *durationBuilder {
	return &durationBuilder{
		params: newDurationParams(),
	}
}


