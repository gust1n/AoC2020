package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readStringsFromFile("./input.txt")

	func() {
		fmt.Println("----part 1----")

		pos := vector{0, 0}
		step := vector{3, 1}
		trees := 0

		for {
			pos = pos.add(step)

			// reached the bottom
			if pos.y >= len(lines) {
				break
			}

			line := lines[pos.y]
			xPos := pos.x % len(line)

			if lines[pos.y][xPos] == '#' {
				trees++
			}

		}

		fmt.Println(trees)
	}()

	func() {
		fmt.Println("----part 2----")

		slopeFn := func(step vector) int {
			pos := vector{0, 0}
			trees := 0

			for {
				pos = pos.add(step)

				// reached the bottom
				if pos.y >= len(lines) {
					break
				}

				line := lines[pos.y]
				xPos := pos.x % len(line)

				if lines[pos.y][xPos] == '#' {
					trees++
				}

			}

			return trees
		}
		first := slopeFn(vector{1, 1})
		second := slopeFn(vector{3, 1})
		third := slopeFn(vector{5, 1})
		fourth := slopeFn(vector{7, 1})
		fifth := slopeFn(vector{1, 2})

		fmt.Println(first * second * third * fourth * fifth)
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

type vector struct {
	x, y int
}

func (v vector) add(other vector) vector {
	return vector{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}
