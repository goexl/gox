package gox

type mapstructOptions struct {
	zeroFields bool
	squash     bool
	tag        string
}

func defaultMapstructOptions() *mapstructOptions {
	return &mapstructOptions{
		zeroFields: true,
		squash:     true,
		tag:        "json",
	}
}
