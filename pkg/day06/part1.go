package day06

import (
	"jordanfitz.com/advent/2024/pkg/util"
)

type Pos struct {
	x, y int
}

func (d day) Part1(inputPath string) any {
	pos := Pos{}

	dir := 0 // 0=up, 1=right, 2=down, 3=left
	result := 1

	var table [][]byte
	for y, row := range util.ReadLines(inputPath) {
		if row == "" {
			continue
		}
		for x, c := range row {
			if c == '^' {
				pos.x = x
				pos.y = y
			}
		}
		table = append(table, []byte(row))
	}

	visited := map[Pos]bool{}

	for {
		var next Pos
		switch dir {
		case 0:
			next.y = pos.y - 1
			next.x = pos.x
		case 1:
			next.y = pos.y
			next.x = pos.x + 1
		case 2:
			next.y = pos.y + 1
			next.x = pos.x
		case 3:
			next.y = pos.y
			next.x = pos.x - 1
		}

		if next.y < 0 || next.x < 0 || next.y >= len(table) || next.x >= len(table[next.y]) {
			break
		}

		if table[next.y][next.x] == '#' {
			dir++
			dir %= 4
		} else {
			pos = next
			if !visited[pos] {
				visited[pos] = true
				result++
			}
		}
	}

	return result
}
