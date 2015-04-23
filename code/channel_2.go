package main

import (
	"fmt"
)

func intGenerator() chan int {
	ch := make(chan int)
	id := 0
	go func() {
		for {
			ch <- id
			id++
		}
	}()
	return ch
}
func main() {
	generator := intGenerator()
	fmt.Println("Int:", <-generator)
	fmt.Println("Int:", <-generator)
	fmt.Println("Int:", <-generator)
}
