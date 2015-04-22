package main

import (
	"log"
)

func multiplicator(n int) func(int) int {
	return func(number int) int {
		return n * number
	}
}

func main() {

	ten := multiplicator(10)
	two := multiplicator(2)

	log.Println("16 times 10:", ten(16))
	log.Println("45 times 2:", two(45))
}
