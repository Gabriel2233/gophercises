package main

import (
	"flag"
	"net/http"
)

func main() {
	urlFlag := flag.String("website", "https://gabriel-tiso-blog.vercel.app", "a website to scan")
	flag.Parse()

	res, err := http.Get(*urlFlag)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
}
