package runner

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"jordanfitz.com/advent/2024/pkg/util"
)

type Parts interface {
	Part1(inputPath string) any
	Part2(inputPath string) any
}

type registeredDay struct {
	parts     Parts
	inputPath string
}

var (
	packagePathRegex = regexp.MustCompile(`jordanfitz.com/advent/20\d{2}/pkg/day\d{2}`)

	nameInputRegex = regexp.MustCompile(`(?i)(d|day)?(\d{1,2})`)
	partInputRegex = regexp.MustCompile(`(?i)(p|part)?(\d{1,2})`)

	registeredDays = make(map[string]registeredDay, 25)
)

func RegisterDay(dayParts Parts) {
	packagePath := reflect.TypeOf(dayParts).PkgPath()
	if !packagePathRegex.Match([]byte(packagePath)) {
		panic(fmt.Sprintf("RegisterDay received bad package format '%s'", packagePath))
	}
	path := strings.Split(packagePath, "/")

	name := path[len(path)-1]
	if _, ok := registeredDays[name]; ok {
		panic(fmt.Sprintf("day '%s' is already registered", name))
	}
	registeredDays[name] = registeredDay{
		parts:     dayParts,
		inputPath: strings.Join(path[len(path)-2:], "/") + "/input.txt",
	}
}

func ExecuteDayPart(name, part string) {
	if groups := nameInputRegex.FindStringSubmatch(name); groups == nil {
		panic(fmt.Sprintf("day '%s' is invalid (bad format)", name))
	} else {
		name = fmt.Sprintf("day%02d", util.Int(groups[2]))
	}
	if groups := partInputRegex.FindStringSubmatch(part); groups == nil {
		panic(fmt.Sprintf("part '%s' is invalid (bad format)", name))
	} else {
		part = fmt.Sprintf("part%d", util.Int(groups[2]))
	}

	day, ok := registeredDays[name]
	if !ok {
		panic(fmt.Sprintf("day '%s' is not registered", name))
	}

	fmt.Printf("Executing %s.%s\n\n", name, part)

	var result any
	switch part {
	case "part1":
		result = day.parts.Part1(day.inputPath)
	case "part2":
		result = day.parts.Part2(day.inputPath)
	default:
		panic(fmt.Sprintf("invalid part specifier '%s'", part))
	}

	fmt.Printf("Result is %v\n", result)
}
