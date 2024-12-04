package main

import (
	"fmt"
	"strings"

	"jordanfitz.com/advent/2024/pkg/util"
)

func isSafe2(nums []int, levelRemoved bool) bool {
	recurse := func() bool {
		for i := range nums {
			n := make([]int, len(nums))
			copy(n, nums)
			n = append(n[:i], n[i+1:]...)
			if isSafe2(n, true) {
				return true
			}
		}
		return false
	}

	var dir *bool
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		if util.IntAbs(diff) > 3 || diff == 0 {
			if levelRemoved {
				return false
			}
			return recurse()
		}
		if dir == nil {
			dir = util.Ptr(diff > 0)
		} else if *dir != (diff > 0) {
			if levelRemoved {
				return false
			}
			return recurse()
		}
	}
	return true
}

func part2() {
	numSafe := 0
	util.ForLines("input.txt", func(report string) {
		if report == "" {
			return
		}
		nums := util.Map(strings.Split(report, " "), func(num string) int {
			return util.Int(num)
		})
		if isSafe2(nums, false) {
			numSafe++
		}
	})
	fmt.Println(numSafe)
}
