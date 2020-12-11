package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

const (
	empty    = 'L'
	floor    = '.'
	occupied = '#'
)

func main() {
	lines := readStringsFromFile("./input.txt")
	surrounding := []vector{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}}

	makeInitialGrid := func() grid {
		grid := map[vector]rune{}
		for rowIdx, row := range lines {
			for colIdx, col := range row {
				v := vector{y: rowIdx, x: colIdx}
				grid[v] = col
			}
		}

		return grid
	}

	numSurroundingOccupied := func(grid grid, v vector, maxStep *int) int {
		var numOccupied int

		for _, dir := range surrounding {
			pos := v
			step := 0
			for maxStep == nil || step < *maxStep {
				nextPos := pos.add(dir)
				next, ok := grid[nextPos]
				if !ok { // reached end of grid
					break
				}
				if next == occupied {
					numOccupied++
					break
				}
				if next == empty {
					break
				}
				step++
				pos = nextPos
			}
		}

		return numOccupied
	}

	processGrid := func(grid grid, strategy func(grid, vector, *int) int, minOccupied int, maxStep *int) grid {
		newGrid := map[vector]rune{}
		for v, status := range grid {
			switch status {
			case empty:
				if strategy(grid, v, maxStep) == 0 {
					newGrid[v] = occupied
				} else {
					newGrid[v] = empty
				}
			case occupied:
				if strategy(grid, v, maxStep) >= minOccupied {
					newGrid[v] = empty
				} else {
					newGrid[v] = occupied
				}
			case floor:
				newGrid[v] = floor
			}
		}

		return newGrid
	}

	countOccupied := func(grid grid) int {
		var numOccupied int
		for _, v := range grid {
			if v == occupied {
				numOccupied++
			}
		}
		return numOccupied
	}

	func() {
		fmt.Println("----part 1----")

		prevGrid := makeInitialGrid()

		for {
			maxStep := 1
			newGrid := processGrid(prevGrid, numSurroundingOccupied, 4, &maxStep)

			if reflect.DeepEqual(newGrid, prevGrid) {
				break
			}

			prevGrid = newGrid
		}

		fmt.Println(countOccupied(prevGrid))
	}()

	func() {
		fmt.Println("----part 2----")

		prevGrid := makeInitialGrid()
		for {
			newGrid := processGrid(prevGrid, numSurroundingOccupied, 5, nil)

			if reflect.DeepEqual(newGrid, prevGrid) {
				break
			}

			prevGrid = newGrid
		}

		fmt.Println(countOccupied(prevGrid))
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

type grid map[vector]rune

type vector struct {
	x, y int
}

func (v vector) add(v2 vector) vector {
	return vector{
		x: v.x + v2.x,
		y: v.y + v2.y,
	}
}
