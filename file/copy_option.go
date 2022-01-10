package file

type (
	copyOption interface {
		applyCopy(options *copyOptions)
	}

	copyOptions struct {
		mode writeMode
	}
)

func defaultCopyOptions() *copyOptions {
	return &copyOptions{
		mode: WriteModeError,
	}
}
