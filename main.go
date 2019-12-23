package hello

import (
	"fmt"
)

var version = "0.1.0"

// Hello data
type Hello struct {
	Name string
}

// SayHello someone with name
func (h *Hello) SayHello() error {
	fmt.Println("Hello baby: ", h.Name)
	return nil
}

// SayHi someone with name
func (h *Hello) SayHi() error {
	fmt.Println("Hi baby: ", h.Name)
	return nil
}

// NewHello create Hello struct
func NewHello(name string) *Hello {
	return &Hello{Name: name}
}
