package day08

import (
	"jordanfitz.com/advent/2024/pkg/util"
)

func walk2(fs []int) {
	fileID := fs[len(fs)-1]

	d := len(fs) - 1
	currentData := func() (start, end int) {
		// advance until we find fileID
		for ; d >= 0 && fs[d] != fileID; d-- {
			if d == 0 {
				return -1, -1
			}
		}
		end = d
		// advance until we find a non-fileID value
		for ; d >= 0 && fs[d] == fileID; d-- {
		}
		start = d + 1
		return start, end
	}

	findBlank := func(dataStart, dataEnd int) (start int) {
		b := 0
		nextEmpty := func() (start, end int) {
			if b == len(fs) {
				return -1, -1
			}
			// advance until we find a -1
			for ; b < len(fs) && fs[b] != -1; b++ {
				if b == len(fs)-1 {
					return -1, -1
				}
			}
			start = b
			// advance until we find a non -1
			for ; b+1 < len(fs) && fs[b+1] == -1; b++ {
				if b+1 == len(fs)-1 {
					return -1, -1
				}
			}
			end = b
			b++
			return start, end
		}

		for s, e := nextEmpty(); s != -1; s, e = nextEmpty() {
			if s >= dataStart {
				return -1
			}
			if e-s >= dataEnd-dataStart {
				return s
			}
		}

		return -1
	}

	for ; fileID >= 0; fileID-- {
		dataStart, dataEnd := currentData()
		blankStart := findBlank(dataStart, dataEnd)
		if blankStart == -1 {
			// we couldn't find a space big enough
			continue
		}
		for i := 0; i < (dataEnd - dataStart + 1); i++ {
			fs[blankStart+i], fs[dataStart+i] = fs[dataStart+i], fs[blankStart+i]
		}
	}
}

func (d day) Part2(inputPath string) any {
	bytes := util.Map([]byte(util.ReadFile(inputPath)), func(b byte) byte {
		return b - 48
	})

	fs := makeFS(bytes)
	walk2(fs)

	return checksum(fs)
}
