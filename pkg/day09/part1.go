package day08

import (
	"fmt"

	"jordanfitz.com/advent/2024/pkg/util"
)

func makeFS(bytes []byte) []int {
	res := make([]int, 0, len(bytes))

	id := 0
	for i, b := range bytes {
		if i%2 == 1 {
			for range b {
				res = append(res, -1)
			}
		} else {
			for range b {
				res = append(res, id)
			}
			id++
		}
	}

	return res
}

func render(fs []int) {
	for _, v := range fs {
		if v == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(v)
		}
	}
	fmt.Println()
}

func walk(fs []int) {
	data := len(fs) - 1
	retreatToData := func() {
		for fs[data] == -1 {
			data--
		}
	}
	retreatToData()

	blank := 0
	advanceToBlank := func() {
		for fs[blank] != -1 {
			blank++
		}
	}
	advanceToBlank()

	for data > blank {
		fs[data], fs[blank] = fs[blank], fs[data]
		retreatToData()
		advanceToBlank()
	}
}

func checksum(fs []int) int {
	res := 0
	for i, v := range fs {
		if v == -1 {
			continue
		}
		res += i * v
	}
	return res
}

func (d day) Part1(inputPath string) any {
	bytes := util.Map([]byte(util.ReadFile(inputPath)), func(b byte) byte {
		return b - 48
	})
	bytes = bytes[:len(bytes)-1]

	fs := makeFS(bytes)
	walk(fs)

	return checksum(fs)
}
