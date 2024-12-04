package main

import (
	"fmt"
	"regexp"

	"jordanfitz.com/advent/2024/pkg/util"
)

func canRead(file, s string, i int) bool {
	if i+len(s) > len(file) {
		return false
	}
	return file[i:i+len(s)] == s
}

func part2() {
	file := util.ReadFile("input.txt")
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	var sections []string

	s := ""
	enabled := true
	for i := 0; i < len(file); i++ {
		if canRead(file, "do()", i) {
			enabled = true
			i += 3
		} else if canRead(file, "don't()", i) {
			if s != "" {
				sections = append(sections, s)
			}
			s = ""
			enabled = false
			i += 6
		} else if enabled {
			s += string(file[i])
		}
	}
	if s != "" {
		sections = append(sections, s)
	}

	fmt.Println(len(sections))

	result := 0
	for _, s := range sections {
		matches := r.FindAllStringSubmatch(s, -1)
		for _, match := range matches {
			result += util.Int(match[1]) * util.Int(match[2])
		}
	}
	fmt.Println(result)
}
