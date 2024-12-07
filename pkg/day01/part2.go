package day01

import (
	"strings"

	"jordanfitz.com/advent/2024/pkg/util"
)

func (d day) Part2(inputPath string) any {
	var left, right []int
	util.ForLines(inputPath, func(line string) {
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
	return result
}
