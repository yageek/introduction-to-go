package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ciphers := [...]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Ciphers:", ciphers)

	fmt.Println("Slice 1:", ciphers[5:])
	fmt.Println("Slice 2:", ciphers[:5])
	fmt.Println("Slice 3:", ciphers[1:5])

	rand.Seed(time.Now().Unix())
	table := make([]int, 10)
	for index, _ := range table {
		table[index] = rand.Intn(100)
	}

	fmt.Println("Table:", table)
}
