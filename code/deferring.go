package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	res, err := http.Get("https://www.openwt.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("HTML:", string(data))
}
