package main

import (
	"log"
	"os"

	. "github.com/zrcoder/cdor"
)

func main() {
	cdor := New().
		Cfg().
		Nodes(
			N("cat").
				Children(
					N("meow", O().F("green")),
				),
			N("dog", O().Sh("circle").L("ddd")),
		).
		Cons(
			C("cat.meow", "dog", O().L("haha").S("red")),
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
