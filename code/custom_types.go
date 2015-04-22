package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) presenting() {
	fmt.Printf("Hello, I'm %s\n", p.Name)
}

func main() {
	a := Person{Name: "Georges"}
	a.presenting()
}
