package d02

import (
	"angelmariages/adventofcode23/utils"
	"fmt"
	"testing"
)

var examplePart2 = `
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

func TestDay02Part02(t *testing.T) {
	sumMul := 0

	utils.IterateOverLines(examplePart2, func(line string) {
		maxReds := getMax(line, "red")
		maxBlues := getMax(line, "blue")
		maxGreens := getMax(line, "green")

		sumMul += maxReds * maxGreens * maxBlues
	})

	fmt.Println("Sum example:", sumMul)

	expected := 2286
	if sumMul != expected {
		t.Logf("Sum of example string should be %d but was %d", expected, sumMul)
		t.FailNow()
	}

	sumMul = 0

	utils.IterateOverFile("./input.txt", func(line string) {
		maxReds := getMax(line, "red")
		maxBlues := getMax(line, "blue")
		maxGreens := getMax(line, "green")

		sumMul += maxReds * maxGreens * maxBlues
	})

	fmt.Println("Sum part2:", sumMul)

	expected = 78375
	if sumMul != expected {
		t.Logf("Sum of example string should be %d but was %d", expected, sumMul)
		t.FailNow()
	}
}
