package main

import "fmt"

// #include <stdio.h>
import "C"

func main() {
	n, err := C.getchar()
	fmt.Print(n, err)
}
