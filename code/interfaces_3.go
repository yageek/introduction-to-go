package main

import (
	"fmt"
)

func main() {
	var unknown interface{} = 2
	_, ok := unknown.(string)
	if ok {
		fmt.Println("Is a string")
	} else {
		fmt.Println("Is not a string :(")
	}
}
