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

	earliestDeparture := parseInt(lines[0])
	var busIDs []int
	for _, s := range strings.Split(lines[1], ",") {
		if s != "x" {
			busIDs = append(busIDs, parseInt(s))
		}
	}

	func() {
		fmt.Println("----part 1----")

		// map next departure -> busID
		nextDepartures := map[int]int{}
		for _, busID := range busIDs {
			// find smallest factor above wanted departure
			factor := math.Ceil(float64(earliestDeparture) / float64(busID))
			departure := busID * int(factor)
			diff := departure - earliestDeparture

			nextDepartures[diff] = busID
		}

		// Find smallest diff
		diffs := make([]int, 0, len(nextDepartures))
		for k := range nextDepartures {
			diffs = append(diffs, k)
		}
		sort.Ints(diffs)

		fmt.Println(diffs[0] * nextDepartures[diffs[0]])
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
