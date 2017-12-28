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
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
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
