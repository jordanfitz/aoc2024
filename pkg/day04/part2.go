package day04

import (
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

func (d day) Part2(inputPath string) any {
	lines := util.ReadLines(inputPath)

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
	return result
}
