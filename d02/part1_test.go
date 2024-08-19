package d02

import (
	"angelmariages/adventofcode23/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var examplePart1 = `
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

type Bag struct {
	red   int
	blue  int
	green int
}

func TestDay02Part01(t *testing.T) {
	exampleMaxes := Bag{
		red:   12,
		green: 13,
		blue:  14,
	}

	sumGameIds := 0

	utils.IterateOverLines(examplePart1, func(line string) {
		gameId := readGameId(line)
		maxReds := getMax(line, "red")
		maxBlues := getMax(line, "blue")
		maxGreens := getMax(line, "green")

		if (isGamePossible(exampleMaxes, Bag{
			red:   maxReds,
			blue:  maxBlues,
			green: maxGreens,
		})) {
			sumGameIds += gameId
		}
	})

	fmt.Println("Sum example:", sumGameIds)

	expected := 8
	if sumGameIds != expected {
		t.Logf("Sum of example string should be %d but was %d", expected, sumGameIds)
		t.FailNow()
	}

	part1Maxes := Bag{
		red:   12,
		green: 13,
		blue:  14,
	}

	sumGameIds = 0

	utils.IterateOverFile("./input.txt", func(line string) {
		gameId := readGameId(line)
		maxReds := getMax(line, "red")
		maxBlues := getMax(line, "blue")
		maxGreens := getMax(line, "green")

		if (isGamePossible(part1Maxes, Bag{
			red:   maxReds,
			blue:  maxBlues,
			green: maxGreens,
		})) {
			sumGameIds += gameId
		}
	})

	fmt.Println("Sum part1:", sumGameIds)

	expected = 2406
	if sumGameIds != expected {
		t.Logf("Sum of example string should be %d but was %d", expected, sumGameIds)
		t.FailNow()
	}
}

func isGamePossible(realBag Bag, gameBag Bag) bool {
	return gameBag.red <= realBag.red && gameBag.blue <= realBag.blue && gameBag.green <= realBag.green
}

func readGameId(input string) int {
	r, _ := regexp.Compile(`Game (\d+):`)

	i, _ := strconv.Atoi(r.FindStringSubmatch(input)[1])

	return i
}

func getMax(input string, color string) int {
	picks := strings.Split(input, ";")

	r, _ := regexp.Compile(`(\d+) ` + color)

	max := 0
	for _, pick := range picks {
		if r.MatchString(pick) {
			newMax, _ := strconv.Atoi(r.FindStringSubmatch(pick)[1])
			if newMax > max {
				max = newMax
			}
		}
	}

	return max
}
