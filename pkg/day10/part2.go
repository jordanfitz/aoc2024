package day08

import (
	"fmt"

	"jordanfitz.com/advent/2024/pkg/util"
)

func numUniqueTrails(mat [][]int, th Pos) int {
	util.Assert(mat[th.y][th.x] == 0)

	trails := map[string]bool{}

	var visit func(p Pos, curr int, trail string)
	visit = func(p Pos, curr int, trail string) {
		if mat[p.y][p.x] == 9 {
			trails[trail] = true
			return
		}
		for _, dir := range directions {
			next := Pos{p.x + dir[0], p.y + dir[1]}
			if next.y < 0 || next.y >= len(mat) || next.x < 0 || next.x >= len(mat[next.y]) {
				continue
			}
			if mat[next.y][next.x] == curr+1 {
				visit(next, curr+1, fmt.Sprintf("%s|%s", trail, next))
			}
		}
	}

	visit(th, 0, th.String())

	return len(trails)
}

func (d day) Part2(inputPath string) any {
	var trailheads []Pos

	y := 0
	mat := util.Map(util.ReadLines(inputPath), func(line string) []int {
		x := 0
		row := util.Map([]byte(line), func(b byte) int {
			val := int(b) - 48
			if val == 0 {
				trailheads = append(trailheads, Pos{x, y})
			}
			x++
			return val
		})
		y++
		return row
	})

	result := 0
	for _, th := range trailheads {
		result += numUniqueTrails(mat, th)
	}

	return result
}
