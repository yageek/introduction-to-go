package main

import "fmt"

func square(n int) int {
	return n * n
}

func rectanglePerimeter(l, L int) int {
	return (l + L) * 2
}

func main() {

	fmt.Println("Square 6:", square(6))
	fmt.Println("Rectangle Perimeter (3,6):", rectanglePerimeter(3, 6))

	someFunction := func(n int) int {
		return n + 4
	}

	fmt.Println("someFunction with 4:", someFunction(4))
}
