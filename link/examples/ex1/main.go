package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gnarula/gophercises/link"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)

	links, err := link.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", links)
}
