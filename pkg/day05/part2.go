package day05

import (
	"slices"
	"strings"

	"jordanfitz.com/advent/2024/pkg/util"
)

func (d day) Part2(inputPath string) any {
	input := util.ReadFile(inputPath)
	chunks := strings.Split(input, "\n\n")
	ordering, updates := strings.TrimSpace(chunks[0]), strings.TrimSpace(chunks[1])

	lookup := map[string]bool{}
	rules := [][2]string{}
	for _, line := range strings.Split(ordering, "\n") {
		lookup[line] = true
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
		if isValid(nums) {
			continue
		}
		slices.SortFunc(nums, func(a, b string) int {
			if lookup[a+"|"+b] {
				return -1
			}
			if lookup[b+"|"+a] {
				return 1
			}
			return 0
		})
		mid := util.Int(nums[len(nums)/2])
		result += mid
	}

	return result
}
