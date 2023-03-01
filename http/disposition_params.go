package http

type dispositionParams struct {
	filename    string
	disposition string
}

func newDispositionParams(filename string) *dispositionParams {
	return &dispositionParams{
		filename:    filename,
		disposition: attachment,
	}
}
