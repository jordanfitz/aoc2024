package main

import (
	"fmt"
	"strconv"
	"strings"

	"jordanfitz.com/advent/2024/pkg/util"
)

func calculate2(operands []int64, operators []Operator) int64 {
	ands := make([]int64, len(operands))
	copy(ands, operands)

	result := operands[0]
	i := 1
	for _, o := range operators {
		switch o {
		case Add:
			result += ands[i]
		case Multiply:
			result *= ands[i]
		case Concat:
			s := util.Str(result) + util.Str(ands[i])
			i64, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				continue
			}
			result = i64
		}
		i++
	}
	return result
}

func check2(expected int64, operands []int64) bool {
	var c func(operators []Operator) bool
	c = func(operators []Operator) bool {
		if len(operators) == len(operands)-1 {
			return calculate2(operands, operators) == expected
		}

		o := make([]Operator, len(operators))
		copy(o, operators)
		if c(append(o, Add)) {
			return true
		}

		o = make([]Operator, len(operators))
		copy(o, operators)
		if c(append(o, Multiply)) {
			return true
		}

		o = make([]Operator, len(operators))
		copy(o, operators)
		return c(append(o, Concat))
	}

	return c(make([]Operator, 0))
}

func part2() {
	result := int64(0)
	util.ForLines("input.txt", func(line string) {
		if line == "" {
			return
		}
		parts := strings.Split(line, ": ")
		expected := util.Int64(parts[0])
		operands := util.Map(strings.Split(parts[1], " "), func(s string) int64 {
			return util.Int64(s)
		})
		if check2(expected, operands) {
			result += expected
		}
	})
	fmt.Println(result)
}
