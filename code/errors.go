package main

import (
	"log"
	"net/http"
)

func main() {

	_, err := http.Get("https://www.openwwdwedwedt.com/")
	if err != nil {
		log.Fatal(err)
	}
}
