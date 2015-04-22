package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(1e6 + 1)
	fmt.Println("i:", i)

	switch {
	case i > 1 && i < 1000:
		fmt.Println("Between one and thousand")
	case i > 1000 && i < 1e6:
		fmt.Println("Between hundred and 1 million")
	default:
		fmt.Println("Too huge ! Lucky guy :)")
	}
}
