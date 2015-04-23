package main

import (
	"fmt"
	"time"
)

func work(ch chan bool, n time.Duration) {
	time.Sleep(n * time.Millisecond)
	fmt.Printf("%d ms\n", n)
	ch <- true
}

func main() {
	ch := make(chan bool)
	go work(ch, 1000)
	go work(ch, 600)
	go work(ch, 300)
	<-ch
	<-ch
	<-ch
	fmt.Println("Finished!")
}
