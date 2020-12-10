package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	ints := readIntsFromFile("./input.txt")
	// add first case
	ints = append([]int{0}, ints...)
	// add last case
	ints = append(ints, max(ints...)+3)
	sort.Ints(ints)

	func() {
		fmt.Println("----part 1----")

		var diffs []int

		for i := 1; i < len(ints); i++ {
			cur := ints[i]
			prev := ints[i-1]
			diffs = append(diffs, cur-prev)
		}

		ones := len(filterInts(diffs, func(i int) bool {
			return i == 1
		}))
		threes := len(filterInts(diffs, func(i int) bool {
			return i == 3
		}))

		fmt.Println(ones * threes)
	}()

	func() {
		fmt.Println("----part 2----")

		currentGroup := []int{ints[0]}
		groups := [][]int{}

		// split in groups with at least 3 between, since we cannot remove anything there
		for i := 1; i < len(ints); i++ {
			current := ints[i]
			prev := ints[i-1]

			if current-prev >= 3 {
				groups = append(groups, currentGroup)
				currentGroup = nil
			}

			currentGroup = append(currentGroup, current)
			prev = current
		}

		combinationsPerGroupSize := map[int]int{
			1: 1,
			2: 1,
			3: 2,
			4: 4,
			5: 7,
		}

		numCombinations := 1
		for _, group := range groups {
			numCombinations = numCombinations * combinationsPerGroupSize[len(group)]
		}

		fmt.Println(numCombinations)
	}()
}

func filterInts(ii []int, test func(int) bool) (ret []int) {
	for _, i := range ii {
		if test(i) {
			ret = append(ret, i)
		}
	}
	return
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

func max(list ...int) int {
	sort.Ints(list)
	return list[len(list)-1]
}
