package d01

import (
	"angelmariages/adventofcode23/utils"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var examplePart2 = `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`

func TestDay01Part02(t *testing.T) {
	sum := 0

	utils.IterateOverLines(examplePart2, func(line string) {
		// fmt.Println("Line:", line, ", Translated:", translateSpelledNumbers(line))

		first, last := getFirstAndLastNumbers2(translateSpelledNumbers(line))

		toSum, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))

		if err != nil {
			panic(err)
		}

		sum += toSum
	})

	fmt.Println("Sum example:", sum)

	expected := 281
	if sum != expected {
		t.Logf("Sum of example string should be %d but was %d", expected, sum)
		t.FailNow()
	}

	sum = 0

	utils.IterateOverFile("./input.txt", func(line string) {
		// fmt.Println("Line:", line, ", Translated:", translateSpelledNumbers(line))

		first, last := getFirstAndLastNumbers2(translateSpelledNumbers(line))

		toSum, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))

		if err != nil {
			panic(err)
		}

		sum += toSum
	})

	fmt.Println("Sum file:", sum)

	expected = 54578
	if sum != expected {
		t.Logf("Sum of input string should be %d but was %d", expected, sum)
		t.FailNow()
	}
}

func translateSpelledNumbers(input string) string {
	spelledNumber, foundAt := getSpelledNumber(input)

	if spelledNumber != -1 {
		spelledNumberLen := len(spelledNumbers[spelledNumber])

		// we need to not replace the whole word
		// as some words overlap and the problem requires to parse both of them
		// so: oneight needs to be 18 and not 1ight
		// or: twone needs to be 21 and not 2ne
		// the trick found in Reddit is to keep first and last letters:
		// oneight -> o1e8t
		// twone -> t2o1e
		input = input[0:foundAt+1] + strconv.Itoa(spelledNumber) + input[spelledNumberLen+foundAt-1:]

		return translateSpelledNumbers(input)
	}

	return input
}

var spelledNumbers = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

func getSpelledNumber(input string) (int, int) {
	return calcSpelledNumber(input, len(input))
}

func calcSpelledNumber(input string, originalInputLength int) (int, int) {
	for k, c := range spelledNumbers {
		if strings.HasPrefix(input, c) {
			return k, originalInputLength - len(input)
		}
	}

	inputLen := len(input)

	if inputLen == 0 {
		return -1, 0
	}

	return calcSpelledNumber(input[1:], originalInputLength)
}

func getFirstAndLastNumbers2(input string) (int, int) {
	first := -1
	last := -1

	for _, r := range input {
		rawI, err := strconv.ParseInt(string(r), 10, 0)

		if err != nil {
			continue
		}

		i := int(rawI)

		if first == -1 {
			first = i
			last = i
		} else {
			last = i
		}
	}

	return first, last
}
