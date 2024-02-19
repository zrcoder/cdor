package main

import (
	"log"
	"os"

	c "github.com/zrcoder/cdor"
)

func main() {
	cdor := c.New().
		Cfg().
		Nodes(
			c.N("cat").
				Sub(
					c.N("meow", c.O().F("green")),
				),
			c.N("dog"),
		).
		Cons(
			c.C("cat.meow", "dog", c.O().L("haha")),
		)

	data, err := cdor.Gen()
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("out.svg", data, 0600)
	if err != nil {
		log.Fatal(err)
	}
}
