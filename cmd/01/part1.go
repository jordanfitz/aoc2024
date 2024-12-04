package main

import (
	"fmt"
	"slices"
	"strings"

	"jordanfitz.com/advent/2024/pkg/util"
)

func part1() {
	var left, right []int
	util.ForLines("input.txt", func(line string) {
		if line == "" {
			return
		}
		parts := strings.Split(line, "   ")
		left = append(left, util.Int(parts[0]))
		right = append(right, util.Int(parts[1]))
	})

	slices.Sort(left)
	slices.Sort(right)

	result := 0
	for i := range left {
		result += util.IntAbs(left[i] - right[i])
	}
	fmt.Println(result)
}
