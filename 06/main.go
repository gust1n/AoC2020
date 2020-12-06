package main

import (
	"bufio"
	"fmt"
	"os"
)

type group struct {
	numPassengers int
	uniqeAnswers  map[rune]int
}

func newGroup() group {
	return group{
		uniqeAnswers: map[rune]int{},
	}
}

func main() {
	lines := readStringsFromFile("./input.txt")
	var groups []group
	current := newGroup()

	for _, line := range lines {
		if line == "" {
			groups = append(groups, current)
			current = newGroup()
		} else {
			current.numPassengers++
		}

		for _, c := range line {
			current.uniqeAnswers[c]++
		}
	}

	func() {
		fmt.Println("----part 1----")

		var count int
		for _, group := range groups {
			count += len(group.uniqeAnswers)
		}

		fmt.Println(count)
	}()

	func() {
		fmt.Println("----part 2----")

		var count int
		for _, group := range groups {
			for _, answers := range group.uniqeAnswers {
				if answers == group.numPassengers {
					count++
				}
			}
		}

		fmt.Println(count)
	}()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readStringsFromFile(path string) []string {
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
