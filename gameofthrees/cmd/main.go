package main

import (
	"fmt"
	"os"
	"strconv"

	"code-katas/gameofthrees"
)

func main() {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("invalid input")
		os.Exit(1)
	}

	for _, n := range gameofthrees.DivideBy3(num) {
		fmt.Println(n)
	}
}
