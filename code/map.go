package main

import (
	"fmt"
)

func main() {

	data := map[string]int{
		"key1": 31,
		"key2": 28,
		"key3": 24,
	}

	for key, value := range data {
		fmt.Printf("Key %s: %d\n", key, value)
	}
}
