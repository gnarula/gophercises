package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gnarula/gophercises/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA application on")
	filename := flag.String("file", "gopher.json", "the JSON file for the CYOA story")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	story, err := cyoa.JsonStory(file)
	if err != nil {
		log.Fatal(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
