package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := readInputLinesFromFile("./input.txt")
	var r = regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): ([a-z]+)$`)

	func() {
		fmt.Println("----part 1----")

		numValid := 0
		for _, line := range lines {
			matches := r.FindStringSubmatch(line)
			min, err := strconv.Atoi(matches[1])
			check(err)
			max, err := strconv.Atoi(matches[2])
			check(err)
			letter := matches[3]
			pw := matches[4]

			num := strings.Count(pw, letter)
			if num >= min && num <= max {
				numValid++
			}

		}
		fmt.Printf("valid passwords: %d", numValid)
	}()

	func() {
		fmt.Printf("\n----part 2----\n")

		numValid := 0
		for _, line := range lines {
			num := 0

			matches := r.FindStringSubmatch(line)
			pos1 := parseInt(matches[1]) - 1
			pos2 := parseInt(matches[2]) - 1
			letter := matches[3][0]
			pw := matches[4]

			if (pw[pos1]) == letter {
				num++
			}
			if (pw[pos2]) == letter {
				num++
			}

			if num == 1 {
				numValid++
			}
		}
		fmt.Printf("valid passwords: %d", numValid)
	}()
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readInputLinesFromFile(path string) []string {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
