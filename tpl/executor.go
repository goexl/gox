package tpl

type executor interface {
	toString(input []string, inputType inputType, data any) (result string, err error)
	toFile(input []string, inputType inputType, data any, filename string) (err error)
}
