package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Gabriel2233/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 4000, "the port in which the server will start")
	filename := flag.String("file", "gopher.json", "specify a json file containing a story (default gopher.json)")
	flag.Parse()

	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)

	if err != nil {
		panic(err)
	}

	handler := cyoa.NewHandler(story)
	fmt.Printf("web app started at: %d\n", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))
}
