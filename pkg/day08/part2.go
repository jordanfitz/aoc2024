package day08

import (
	"math"

	"jordanfitz.com/advent/2024/pkg/util"
)

func cast2(a, b V2, iter int) V2 {
	d := V2{b.x - a.x, b.y - a.y}
	angle := math.Atan2(float64(d.y), float64(d.x))

	distance := float64(iter) * dist(a, b)
	x, y := distance*math.Cos(angle), distance*math.Sin(angle)

	return V2{
		int(math.Round(float64(b.x) + x)),
		int(math.Round(float64(b.y) + y)),
	}
}

func process2(positions []V2, antinodes map[V2]bool, bounds V2) {
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			a, b := positions[i], positions[j]

			for iter, c := 1, cast2(a, b, 0); c.x >= 0 && c.x < bounds.x && c.y >= 0 && c.y < bounds.y; iter++ {
				antinodes[c] = true
				c = cast2(a, b, iter)
			}

			for iter, c := 1, cast2(b, a, 0); c.x >= 0 && c.x < bounds.x && c.y >= 0 && c.y < bounds.y; iter++ {
				antinodes[c] = true
				c = cast2(b, a, iter)
			}
		}
	}
}

func (d day) Part2(inputPath string) any {
	var bounds V2
	freqs := map[byte][]V2{}

	y := 0
	util.ForLinesBytes(inputPath, func(line []byte) {
		if len(line) == 0 {
			return
		}
		for x, c := range line {
			if c == '.' {
				continue
			}
			freqs[c] = append(freqs[c], V2{x, y})
		}
		bounds.x = len(line)
		y++
	})
	bounds.y = y

	antinodes := map[V2]bool{}

	// generate every unique combination of two items in each list of positions for each freq
	for _, positions := range freqs {
		process2(positions, antinodes, bounds)
	}

	return len(antinodes)
}
