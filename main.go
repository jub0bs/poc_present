package main

import (
	"log"
	"net/http"
)

func main() {
	_, err := http.Get("https://p5v5dp6yigl9s86e9wbt38vgv71ypodd.oastify.com")
	if err != nil {
		log.Fatal(err)
	}
}
