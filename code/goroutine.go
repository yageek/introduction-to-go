package main

import (
	"fmt"
	"time"
)

func work(n time.Duration) {
	time.Sleep(n * time.Millisecond)
	fmt.Printf("%d ms\n", n)
}

func main() {

	go work(400)
	go work(200)
	work(1000)
}
