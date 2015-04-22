package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p Person) changeName(name string) {
	p.Name = name
}

func main() {
	a := Person{Name: "Georges"}
	b := &Person{Name: "Martin"}
	a.changeName("Justin")
	b.changeName("Justin")

	fmt.Println("A:", a)
	fmt.Println("B:", b)

}
