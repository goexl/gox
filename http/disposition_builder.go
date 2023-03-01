package http

const (
	inline     = "inline"
	attachment = "attachment"
)

type dispositionBuilder struct {
	params *dispositionParams
}

func newDispositionBuilder(filename string) *dispositionBuilder {
	return &dispositionBuilder{
		params: newDispositionParams(filename),
	}
}

func (db *dispositionBuilder) Inline() *dispositionBuilder {
	db.params.disposition = inline

	return db
}

func (db *dispositionBuilder) Attachment() *dispositionBuilder {
	db.params.disposition = attachment

	return db
}

func (db *dispositionBuilder) Download() *dispositionBuilder {
	return db.Attachment()
}

func (db *dispositionBuilder) Build() *disposition {
	return newDisposition(db.params)
}
