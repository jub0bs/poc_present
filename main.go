package main

import (
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://169.254.169.254/latest/meta-data")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	dump, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(dump)
}
