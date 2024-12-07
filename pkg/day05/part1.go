package day05

import (
	"slices"
	"strings"

	"jordanfitz.com/advent/2024/pkg/util"
)

func (d day) Part1(inputPath string) any {
	input := util.ReadFile(inputPath)
	chunks := strings.Split(input, "\n\n")
	ordering, updates := strings.TrimSpace(chunks[0]), strings.TrimSpace(chunks[1])

	rules := [][2]string{}
	for _, line := range strings.Split(ordering, "\n") {
		parts := strings.Split(line, "|")
		rules = append(rules, [2]string{parts[0], parts[1]})
	}

	isValid := func(nums []string) bool {
		for _, rule := range rules {
			a := slices.Index(nums, rule[0])
			if a == -1 {
				continue
			}
			b := slices.Index(nums, rule[1])
			if b == -1 {
				continue
			}
			if a >= b {
				return false
			}
		}
		return true
	}

	result := 0
	for _, line := range strings.Split(updates, "\n") {
		nums := strings.Split(line, ",")
		if !isValid(nums) {
			continue
		}
		mid := util.Int(nums[len(nums)/2])
		result += mid
	}

	return result
}
