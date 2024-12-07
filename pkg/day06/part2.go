package day06

import (
	"jordanfitz.com/advent/2024/pkg/util"
)

// returns whether it's a loop
func walk(table [][]byte, pos, obs Pos) bool {
	type Loc struct {
		x, y, dir int
	}
	visited := map[Loc]bool{}

	dir := 0
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

		if table[next.y][next.x] == '#' || next.y == obs.y && next.x == obs.x {
			dir++
			dir %= 4
		} else {
			pos = next

			l := Loc{pos.x, pos.y, dir}
			if !visited[l] {
				visited[l] = true
			} else {
				return true
			}
		}
	}

	return false
}

func (d day) Part2(inputPath string) any {
	pos := Pos{}

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

	result := 0
	for y := range table {
		for x := range table[y] {
			if walk(table, pos, Pos{x, y}) {
				result++
			}
		}
	}
	return result
}
