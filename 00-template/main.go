package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines := readStringsFromFile("./input.txt")

	func() {
		fmt.Println("----part 1----")

		fmt.Println("results")
	}()

	func() {
		fmt.Println("----part 2----")

		fmt.Println("results")
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
		y: v.y + v2.y,
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
	m := list[0]
	for _, i := range list {
		if i < m {
			m = i
		}
	}

	return m
}

func max(list ...int) int {
	m := list[0]
	for _, i := range list {
		if i > m {
			m = i
		}
	}

	return m
}
