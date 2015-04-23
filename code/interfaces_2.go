package main

import (
	"fmt"
)

type Count int

func main() {

	values := []interface{}{1, "fwefwefw", Count(9)}
	for _, value := range values {

		switch value.(type) {
		case int:
			fmt.Printf("Value is a int %d\n", value)
		case string:
			fmt.Printf("Is a string:%s\n", value)
		case Count:
			fmt.Printf("Is a Count:%d\n", value)
		}
	}
}
