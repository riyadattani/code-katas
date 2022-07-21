package main

import (
	"fmt"
	"log"
	"os"

	"code-katas/wordwrap"
)

func main() {
	input := os.Args[1]

	if len(os.Args) != 2 {
		log.Fatalf("please add some text to wrap, thaaaas")
	}

	fmt.Printf("%s", wordwrap.Wrap(input, 59))
}
