package main

import (
	"fmt"

	"jordanfitz.com/advent/2024/pkg/util"
)

func find(mat []string, x, y int, char byte) bool {
	if x < 0 || y < 0 || y >= len(mat) || x >= len(mat[y]) {
		return false
	}
	return mat[y][x] == char
}

func search(mat []string, x, y int) int {
	result := 0

	// horiz/vert
	if find(mat, x+1, y, 'M') &&
		find(mat, x+2, y, 'A') &&
		find(mat, x+3, y, 'S') {
		result++
	}
	if find(mat, x-1, y, 'M') &&
		find(mat, x-2, y, 'A') &&
		find(mat, x-3, y, 'S') {
		result++
	}
	if find(mat, x, y+1, 'M') &&
		find(mat, x, y+2, 'A') &&
		find(mat, x, y+3, 'S') {
		result++
	}
	if find(mat, x, y-1, 'M') &&
		find(mat, x, y-2, 'A') &&
		find(mat, x, y-3, 'S') {
		result++
	}

	// diags
	if find(mat, x+1, y+1, 'M') &&
		find(mat, x+2, y+2, 'A') &&
		find(mat, x+3, y+3, 'S') {
		result++
	}
	if find(mat, x-1, y+1, 'M') &&
		find(mat, x-2, y+2, 'A') &&
		find(mat, x-3, y+3, 'S') {
		result++
	}
	if find(mat, x+1, y-1, 'M') &&
		find(mat, x+2, y-2, 'A') &&
		find(mat, x+3, y-3, 'S') {
		result++
	}
	if find(mat, x-1, y-1, 'M') &&
		find(mat, x-2, y-2, 'A') &&
		find(mat, x-3, y-3, 'S') {
		result++
	}

	return result
}

func part1() {
	lines := util.ReadLines("input.txt")

	result := 0
	for y, line := range lines {
		if line == "" {
			continue
		}
		for x, c := range line {
			if c == 'X' {
				result += search(lines, x, y)
			}
		}
	}
	fmt.Println(result)
}
