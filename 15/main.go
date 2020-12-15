package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := readStringsFromFile("./input.txt")
	var initialNumbers []int
	for _, s := range strings.Split(lines[0], ",") {
		initialNumbers = append(initialNumbers, parseInt(s))
	}

	run := func(end int) int {
		spoken := map[int]int{} // store all numbers and the turn they were last spoken

		// fill with initial numbers
		for i, n := range initialNumbers {
			spoken[n] = i + 1
		}

		lastSpoken := initialNumbers[len(initialNumbers)-1]

		for turn := len(initialNumbers); turn < end; turn++ {
			var next int // default 0
			if n, ok := spoken[lastSpoken]; ok {
				next = turn - n
			}
			spoken[lastSpoken] = turn
			lastSpoken = next
		}

		return lastSpoken
	}

	func() {
		fmt.Println("----part 1----")

		fmt.Println(run(2020))

	}()

	func() {
		fmt.Println("----part 2----")

		fmt.Println(run(30000000))
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

func readIntsFromFile(path string) []int {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var ints []int

	for scanner.Scan() {
		ints = append(ints, parseInt(scanner.Text()))
	}

	return ints
}

type vector struct {
	x, y int
}

func (v vector) add(v2 vector) vector {
	return vector{
		x: v.x + v2.x,
		y: v.y + v2.y,
	}
}

func (v vector) sub(v2 vector) vector {
	return vector{
		x: v.x - v2.x,
		y: v.y - v2.y,
	}
}

func (v vector) manhattanDistance(v2 vector) int {
	v = v.sub(v2)

	return abs(v.x + v.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(list ...int) int {
	sort.Ints(list)
	return list[0]
}

func max(list ...int) int {
	sort.Ints(list)
	return list[len(list)-1]
}
