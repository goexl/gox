package tpl

type executor interface {
	toString(template string, data any) (result string, err error)
	toFile(template string, data any, filename string) (err error)
}
