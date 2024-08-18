package utils

import (
	"bufio"
	"os"
	"strings"
)

func IterateOverFile(file string, cb func(line string)) {
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()

		if len(strings.TrimSpace(text)) == 0 {
			continue
		}

		cb(text)
	}
}

func IterateOverLines(lines string, cb func(line string)) {
	scanner := bufio.NewScanner(strings.NewReader(lines))

	for scanner.Scan() {
		text := scanner.Text()

		if len(strings.TrimSpace(text)) == 0 {
			continue
		}

		cb(text)
	}
}
