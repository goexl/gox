package file

type (
	walkOption interface {
		applyWalk(options *walkOptions)
	}

	walkOptions struct {
		pattern   string
		matchable matchable
	}
)

func defaultWalkOptions() *walkOptions {
	return &walkOptions{
		pattern: `*`,
	}
}
