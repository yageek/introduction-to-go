package main

import (
	"fmt"
	"module"
)

func main() {
	//OK
	fmt.Println("Square 3", module.Square(3))

	//ERROR
	fmt.Println("pow 3", module.pow3(3))
}
