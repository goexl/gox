package gox

// Value 值
type Value interface {
	// Value 具体的值
	Value() interface{}
}
