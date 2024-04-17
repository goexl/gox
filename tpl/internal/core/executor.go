package core

type Executor interface {
	ToString(input []string, inputType InputType, name string, data any) (result string, err error)
	ToFile(input []string, inputType InputType, name string, data any, filename string) (err error)
}
