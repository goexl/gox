package tpl

type builder struct {
	data        any
	template    string
	contentType contentType
}

func newBuilder(template string) *builder {
	return &builder{
		template:    template,
		contentType: contentTypeString,
	}
}

func (b *builder) Data(data any) *builder {
	b.data = data

	return b
}

func (b *builder) File() *builder {
	b.contentType = contentTypeFile

	return b
}

func (b *builder) String() *builder {
	b.contentType = contentTypeString

	return b
}
