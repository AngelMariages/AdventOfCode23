package d01

import (
	"angelmariages/adventofcode23/utils"
	"fmt"
	"strconv"
	"testing"
)

var examplePart1 = `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

func TestPart1(t *testing.T) {
	sum := 0

	utils.IterateOverLines(examplePart1, func(line string) {
		first, last := getFirstAndLastNumbers(line)

		toSum, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))

		if err != nil {
			panic(err)
		}

		sum += toSum
	})

	fmt.Println("Sum example:", sum)

	if sum != 142 {
		t.Logf("Sum of example string should be 142 but was %d", sum)
		t.FailNow()
	}

	sum = 0

	utils.IterateOverFile("./input.txt", func(line string) {
		first, last := getFirstAndLastNumbers(line)

		toSum, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))

		if err != nil {
			panic(err)
		}

		sum += toSum
	})

	fmt.Println("Sum file:", sum)

	if sum != 55208 {
		t.Logf("Sum of input string should be 55208 but was %d", sum)
		t.FailNow()
	}
}

func getFirstAndLastNumbers(input string) (int, int) {
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
