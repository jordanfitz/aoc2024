package main

import (
	"fmt"
	"strings"

	"jordanfitz.com/advent/2024/pkg/util"
)

func part2() {
	var left, right []int
	util.ForLines("input.txt", func(line string) {
		if line == "" {
			return
		}
		parts := strings.Split(line, "   ")
		left = append(left, util.Int(parts[0]))
		right = append(right, util.Int(parts[1]))
	})

	occurrences := map[int]int{}
	for _, val := range right {
		occurrences[val]++
	}

	result := 0
	for _, val := range left {
		result += val * occurrences[val]
	}
	fmt.Println(result)
}
