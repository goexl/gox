package tpl

type executor interface {
	toString(input []string, inputType inputType, name string, data any) (result string, err error)
	toFile(input []string, inputType inputType, name string, data any, filename string) (err error)
}
