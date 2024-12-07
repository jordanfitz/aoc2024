package day02

import (
	"strings"

	"jordanfitz.com/advent/2024/pkg/util"
)

func isSafe(nums []int) bool {
	var dir *bool
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		if util.IntAbs(diff) > 3 || diff == 0 {
			return false
		}
		if dir == nil {
			dir = util.Ptr(diff > 0)
		} else if *dir != (diff > 0) {
			return false
		}
	}
	return true
}

func (d day) Part1(inputPath string) any {
	numSafe := 0
	util.ForLines(inputPath, func(report string) {
		if report == "" {
			return
		}
		nums := util.Map(strings.Split(report, " "), func(num string) int {
			return util.Int(num)
		})
		if isSafe(nums) {
			numSafe++
		}
	})
	return numSafe
}
