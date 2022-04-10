package oop

type A struct {
	Name string
}

type B struct {
	Name string
	A
}

func t() {
	b := new(B)
	b.Name = "sda"
}
