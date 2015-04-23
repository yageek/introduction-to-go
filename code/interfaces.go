package main

import "fmt"

type Animal interface {
	Noise() string
}

type Dog struct{}

func (d *Dog) Noise() string {
	return "waf"
}

type Cat struct{}

func (c *Cat) Noise() string {
	return "miaou"
}

func main() {
	var animalA Animal = &Dog{}
	var animalB Animal = &Cat{}
	fmt.Println("Animal A:", animalA.Noise())
	fmt.Println("Animal B:", animalB.Noise())
}
