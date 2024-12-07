package day03

import (
	"regexp"

	"jordanfitz.com/advent/2024/pkg/util"
)

func (d day) Part1(inputPath string) any {
	file := util.ReadFile(inputPath)
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(file, -1)
	result := 0
	for _, match := range matches {
		result += util.Int(match[1]) * util.Int(match[2])
	}
	return result
}
