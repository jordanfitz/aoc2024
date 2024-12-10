package main

import (
	"fmt"
	"os"

	"jordanfitz.com/advent/2024/pkg/runner"

	_ "jordanfitz.com/advent/2024/pkg/days"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: main.go <day> <part>\n")
		os.Exit(1)
	}
	runner.ExecuteDayPart(args[0], args[1])
}
