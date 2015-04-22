package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(20)

	for index := 0; index < i; index++ {
		fmt.Println("Index:", index)
	}
}
