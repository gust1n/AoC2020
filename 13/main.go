package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := readStringsFromFile("./input.txt")

	func() {
		fmt.Println("----part 1----")

		earliestDeparture := parseInt(lines[0])
		var busIDs []int
		for _, s := range strings.Split(lines[1], ",") {
			if s != "x" {
				busIDs = append(busIDs, parseInt(s))
			}
		}

		smallestDiff := math.MaxInt64
		smallestDiffBusID := -1
		for _, busID := range busIDs {
			diff := busID - (earliestDeparture % busID)
			if diff < smallestDiff {
				smallestDiff = diff
				smallestDiffBusID = busID
			}
		}

		fmt.Println(smallestDiff * smallestDiffBusID)
	}()

	func() {
		fmt.Println("----part 2----")

		// map offset -> busID
		busOffsets := map[int]int{}
		for offset, s := range strings.Split(lines[1], ",") {
			if s != "x" {
				busOffsets[offset] = parseInt(s)
			}
		}

		factor := 1
		departureOfFirstBus := 1
		// find factor that matches each bus
		for offset, busID := range busOffsets {
			var multiplier int
			for {
				nextTimestamp := departureOfFirstBus + multiplier*factor
				nextDeparture := nextTimestamp + offset
				if nextDeparture%busID == 0 { // if in time table
					factor *= busID
					departureOfFirstBus = nextTimestamp
					fmt.Println("found factor", factor, offset)
					break
				}
				multiplier++
			}
		}

		fmt.Println(departureOfFirstBus)

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

func max(list ...int) int {
	sort.Ints(list)
	return list[len(list)-1]
}
