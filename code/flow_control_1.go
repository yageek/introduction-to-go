package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(1e7)
	fmt.Println("i:", i)

	if i%2 == 0 {
		fmt.Println("i even")
	} else {
		fmt.Println("i is odd")
	}
}
