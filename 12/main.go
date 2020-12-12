package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var instructionRegexp = regexp.MustCompile(`^([NSEWLRF])(\d+)$`)

type instruction struct {
	op    string
	value int
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

func (v vector) times(m int) vector {
	return vector{
		x: v.x * m,
		y: v.y * m,
	}
}

func main() {
	lines := readStringsFromFile("./input.txt")

	north := vector{0, -1}
	east := vector{1, 0}
	south := vector{0, 1}
	west := vector{-1, 0}

	directions := []vector{north, east, south, west}

	var program []instruction
	for _, line := range lines {
		matches := instructionRegexp.FindStringSubmatch(line)
		program = append(program, instruction{op: matches[1], value: parseInt(matches[2])})
	}

	func() {
		fmt.Println("----part 1----")

		facingIdx := 1

		var currentPos vector

		for _, instruction := range program {
			var currentDirection vector

			switch instruction.op {
			case "F":
				currentDirection = directions[facingIdx]
			case "N":
				currentDirection = north
			case "S":
				currentDirection = south
			case "E":
				currentDirection = east
			case "W":
				currentDirection = west
			case "R":
				steps := instruction.value / 90
				facingIdx = (facingIdx + steps) % len(directions)
				continue
			case "L":
				steps := instruction.value / 90
				facingIdx = (facingIdx - steps)
				if facingIdx < 0 {
					facingIdx = len(directions) + facingIdx
				}
				continue
			}

			for i := 0; i < instruction.value; i++ {
				currentPos = currentPos.add(currentDirection)
			}
		}

		fmt.Println(currentPos.x + currentPos.y)
	}()

	func() {
		fmt.Println("----part 2----")

		waypointDiff := vector{10, -1}
		var currentPos vector

		for _, instruction := range program {
			switch instruction.op {
			case "F":
				step := waypointDiff.times(instruction.value)
				currentPos = currentPos.add(step)
			case "N":
				waypointDiff = waypointDiff.add(north.times(instruction.value))
			case "S":
				waypointDiff = waypointDiff.add(south.times(instruction.value))
			case "E":
				waypointDiff = waypointDiff.add(east.times(instruction.value))
			case "W":
				waypointDiff = waypointDiff.add(west.times(instruction.value))
			case "R":
				for i := 0; i < (instruction.value / 90); i++ {
					waypointDiff = vector{x: -waypointDiff.y, y: waypointDiff.x}
				}
			case "L":
				for i := 0; i < (instruction.value / 90); i++ {
					waypointDiff = vector{x: waypointDiff.y, y: -waypointDiff.x}
				}
			}
		}

		fmt.Println(abs(currentPos.x + currentPos.y))
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
