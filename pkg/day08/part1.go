package day08

import (
	"math"

	"jordanfitz.com/advent/2024/pkg/util"
)

type V2 struct {
	x, y int
}

func dist(a, b V2) float64 {
	x1, y1 := float64(a.x), float64(a.y)
	x2, y2 := float64(b.x), float64(b.y)
	return math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
}

func cast(a, b V2) V2 {
	d := V2{b.x - a.x, b.y - a.y}
	angle := math.Atan2(float64(d.y), float64(d.x))

	distance := dist(a, b)
	x, y := distance*math.Cos(angle), distance*math.Sin(angle)

	return V2{
		int(math.Round(float64(b.x) + x)),
		int(math.Round(float64(b.y) + y)),
	}
}

func process(positions []V2, antinodes map[V2]bool, bounds V2) {
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			a, b := positions[i], positions[j]
			if c := cast(a, b); c.x >= 0 && c.x < bounds.x && c.y >= 0 && c.y < bounds.y {
				antinodes[c] = true
			}
			if c := cast(b, a); c.x >= 0 && c.x < bounds.x && c.y >= 0 && c.y < bounds.y {
				antinodes[c] = true
			}
		}
	}
}

func (d day) Part1(inputPath string) any {
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

	// an antinode occurs at any point that is perfectly in line with two antennas of the same frequency -
	// but only when one of the antennas is twice as far away as the other.

	// point along the line formed by any two antennas of the same freq at which one of the two is twice
	// as far away as the other from said point
	antinodes := map[V2]bool{}

	// generate every unique combination of two items in each list of positions for each freq
	for _, positions := range freqs {
		process(positions, antinodes, bounds)
	}

	return len(antinodes)
}
