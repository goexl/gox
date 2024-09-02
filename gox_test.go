package gox_test

type User struct {
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	Age  int    `json:"age,omitempty" xml:"age,omitempty"`
}
