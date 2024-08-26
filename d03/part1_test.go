package d02

import (
	"angelmariages/adventofcode23/utils"
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

var examplePart1 = `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`

func TestDay03Algo(t *testing.T) {
	if !isNextToSymbolSameLine(4, 5, "....7+....") {
		t.Logf("Is next to symbol line: [%s]", "....7+....")
		t.Fail()
	}

	if !isNextToSymbolSameLine(5, 6, "....+7") {
		t.Logf("Is next to symbol line: [%s]", "....+7")
		t.Fail()
	}
}

func TestDay03Part01(t *testing.T) {
	var lines []string

	utils.IterateOverLines(examplePart1, func(line string) {
		lines = append(lines, line)
	})

	sum := 0

	numberRegex := regexp.MustCompile(`\d+`)
	for lineIdx, line := range lines {
		numbers := numberRegex.FindAllStringIndex(line, -1)
		for _, numIdxs := range numbers {
			// fmt.Println(numIdxs, line[numIdxs[0]:numIdxs[1]])
			if isNextToSymbolSameLine(numIdxs[0], numIdxs[1], line) {
				parsed, _ := strconv.ParseInt(string(line[numIdxs[0]:numIdxs[1]]), 10, 16)

				sum += int(parsed)
			} else if isNextToSymbolVertically(numIdxs[0], numIdxs[1], lines, lineIdx) {
				parsed, _ := strconv.ParseInt(string(line[numIdxs[0]:numIdxs[1]]), 10, 16)

				sum += int(parsed)
			}
		}
	}

	fmt.Println("Sum example:", sum)

	expected := 4361
	if sum != expected {
		t.Logf("Sum of example string should be %d but was %d", expected, sum)
		t.FailNow()
	}

	sum = 0

	lines = nil
	utils.IterateOverFile("./input.txt", func(line string) {
		lines = append(lines, line)
	})

	for lineIdx, line := range lines {
		numbers := numberRegex.FindAllStringIndex(line, -1)
		for _, numIdxs := range numbers {
			// fmt.Println(numIdxs, line[numIdxs[0]:numIdxs[1]])
			if isNextToSymbolSameLine(numIdxs[0], numIdxs[1], line) {
				parsed, _ := strconv.ParseInt(string(line[numIdxs[0]:numIdxs[1]]), 10, 16)

				sum += int(parsed)
			} else if isNextToSymbolVertically(numIdxs[0], numIdxs[1], lines, lineIdx) {
				parsed, _ := strconv.ParseInt(string(line[numIdxs[0]:numIdxs[1]]), 10, 16)

				sum += int(parsed)
			}
		}
	}

	fmt.Println("Sum part1:", sum)

	expected = 522726
	if sum != expected {
		t.Logf("Sum of part1 input should be %d but was %d", expected, sum)
		t.FailNow()
	}
}

func isNextToSymbolSameLine(startIdx, endIdx int, line string) bool {
	r := regexp.MustCompile(`\.+`)

	if startIdx != 0 {
		if !r.MatchString(string(line[startIdx-1])) {
			return true
		}
	}

	if endIdx < len(line)-1 {
		if !r.MatchString(string(line[endIdx])) {
			return true
		}
	}

	return false
}

func isNextToSymbolVertically(startIdx, endIdx int, lines []string, lineIdx int) bool {
	r := regexp.MustCompile(`[\.\d]+`)

	startIdxWithDiagonal := startIdx
	if startIdx > 0 {
		startIdxWithDiagonal = startIdx - 1
	}
	endIdxWithDiagonal := endIdx
	if endIdx < len(lines)-1 {
		endIdxWithDiagonal = endIdx + 1
	}

	for i := startIdxWithDiagonal; i < endIdxWithDiagonal; i++ {
		prevLineIdx := lineIdx - 1
		if lineIdx > 0 && !r.MatchString(string(lines[prevLineIdx][i])) {
			return true
		}

		nextLineIdx := lineIdx + 1
		if lineIdx < len(lines)-1 && !r.MatchString(string(lines[nextLineIdx][i])) {
			return true
		}
	}
	return false
}
