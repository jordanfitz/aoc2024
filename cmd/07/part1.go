package main

import (
	"fmt"
	"strings"

	"jordanfitz.com/advent/2024/pkg/util"
)

type Operator byte

const (
	Add Operator = iota
	Multiply
	Concat
)

func calculate(operands []int64, operators []Operator) int64 {
	result := operands[0]
	i := 1
	for _, o := range operators {
		switch o {
		case Add:
			result += operands[i]
		case Multiply:
			result *= operands[i]
		default:
			panic("wtf")
		}
		i++
	}
	return result
}

func check(expected int64, operands []int64) bool {
	var c func(operators []Operator) bool
	c = func(operators []Operator) bool {
		if len(operators) == len(operands)-1 {
			return calculate(operands, operators) == 64
		}

		o := make([]Operator, len(operators))
		copy(o, operators)
		if c(append(o, Add)) {
			return true
		}

		o = make([]Operator, len(operators))
		copy(o, operators)
		return c(append(o, Multiply))
	}

	return c(make([]Operator, 0))
}

func part1() {
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
		if check(expected, operands) {
			result += expected
		}
	})
	fmt.Println(result)
}
