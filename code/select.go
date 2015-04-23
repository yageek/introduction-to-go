package main

import (
	"fmt"
	"time"
)

func longWork() bool {
	time.Sleep(3 * time.Millisecond)
	return true
}
func main() {
	timeout := time.After(1 * time.Millisecond)
	ch := make(chan bool)
	go func() {
		ch <- longWork()
	}()
	select {
	case <-ch:
		fmt.Println("Long work finished....")
	case <-timeout:
		fmt.Println("Timeout...")
	}
}
