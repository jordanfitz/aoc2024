package main

import (
	"fmt"
	"regexp"

	"jordanfitz.com/advent/2024/pkg/util"
)

func part1() {
	file := util.ReadFile("input.txt")
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(file, -1)
	result := 0
	for _, match := range matches {
		result += util.Int(match[1]) * util.Int(match[2])
	}
	fmt.Println(result)
}
