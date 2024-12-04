package main

import (
	"fmt"

	"jordanfitz.com/advent/2024/pkg/util"
)

func search2(mat []string, x, y int) int {
	result := 0

	if (find(mat, x+1, y+1, 'M') && find(mat, x-1, y-1, 'S')) ||
		(find(mat, x+1, y+1, 'S') && find(mat, x-1, y-1, 'M')) {
		if (find(mat, x+1, y-1, 'M') && find(mat, x-1, y+1, 'S')) ||
			(find(mat, x+1, y-1, 'S') && find(mat, x-1, y+1, 'M')) {
			result++
		}
	}

	return result
}

func part2() {
	lines := util.ReadLines("input.txt")

	result := 0
	for y, line := range lines {
		if line == "" {
			continue
		}
		for x, c := range line {
			if c == 'A' {
				result += search2(lines, x, y)
			}
		}
	}
	fmt.Println(result)
}
