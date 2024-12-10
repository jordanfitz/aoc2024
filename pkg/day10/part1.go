package day08

import (
	"fmt"

	"jordanfitz.com/advent/2024/pkg/util"
)

type Pos struct {
	x, y int
}

func (p Pos) String() string {
	return fmt.Sprint("%d,%d", p.x, p.y)
}

var directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func numUniqueEndpoints(mat [][]int, start Pos) int {
	util.Assert(mat[start.y][start.x] == 0)

	visited := map[Pos]bool{}
	result := 0

	var visit func(p Pos, curr int)
	visit = func(p Pos, curr int) {
		visited[p] = true
		if mat[p.y][p.x] == 9 {
			result++
			return
		}
		for _, dir := range directions {
			next := Pos{p.x + dir[0], p.y + dir[1]}
			if next.x < 0 || next.x >= len(mat[0]) || next.y < 0 || next.y >= len(mat) {
				continue
			}
			if !visited[next] && mat[next.y][next.x] == curr+1 {
				visit(next, curr+1)
			}
		}
	}

	visit(start, 0)

	return result
}

func (d day) Part1(inputPath string) any {
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
	_ = mat

	result := 0
	for _, th := range trailheads {
		result += numUniqueEndpoints(mat, th)
	}

	return result
}
