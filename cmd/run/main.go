package main

import (
	"flag"
	"fmt"
	"os"

	_ "jordanfitz.com/advent/2024/pkg/days"
	"jordanfitz.com/advent/2024/pkg/runner"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: main.go <day> <part>\n\nflags:\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	var (
		positional,
		flags []string
		inputName string
	)

	// hacky solution to allow flags to be inserted anyway around the positional arguments
	rawArgs := os.Args[1:]
	for _, arg := range rawArgs {
		if arg == "" {
			continue
		}
		if arg[0] == '-' {
			flags = append(flags, arg)
		} else {
			positional = append(positional, arg)
		}
	}

	os.Args = append(os.Args[:1], flags...)
	flag.Usage = usage
	flag.StringVar(&inputName, "i", "input.txt", "the input file name")
	flag.Parse()

	if inputName == "" {
		flag.Usage()
	}

	runner.ExecuteDayPart(positional[0], positional[1], inputName)
}
