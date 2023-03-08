package args

type creator struct {
	params *params
}

func newCreator() *creator {
	return &creator{
		params: newParams(),
	}
}

func (c *creator) Capacity(capacity int) *creator {
	c.params.capacity = capacity

	return c
}

func (c *creator) Short(placeholder string) *creator {
	c.params.short = placeholder

	return c
}

func (c *creator) Long(placeholder string) *creator {
	c.params.long = placeholder

	return c
}

func (c *creator) Equal(placeholder string) *creator {
	c.params.equal = placeholder

	return c
}

func (c *creator) Build() *Builder {
	return newBuilder(c.params)
}
