package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	ints := readIntsFromFile("./input.txt")
	var part1results int

	func() {
		fmt.Println("----part 1----")

		preambleSize := 25

		for i := preambleSize; i < len(ints); i++ {
			wanted := ints[i]
			nums := ints[i-preambleSize : i]
			var found bool

			// check each num agains each others
		LOOP:
			for _, n1 := range nums {
				for _, n2 := range nums {
					if n1+n2 == wanted {
						found = true
						break LOOP
					}
				}
			}

			if !found {
				part1results = wanted
			}
		}

		fmt.Println(part1results)
	}()

	func() {
		fmt.Println("----part 2----")

		wantedSetSum := part1results

	LOOP:
		for setFrom := range ints {
			// grow set until either found a match, or reached end
			for setTo := setFrom + 2; setTo < len(ints); setTo++ {
				currentSet := ints[setFrom:setTo]
				currentSetSum := 0

				for _, n := range currentSet {
					currentSetSum += n
					if currentSetSum == wantedSetSum {
						// sort.Ints(currentSet)
						max, min := max(currentSet...), min(currentSet...)
						fmt.Println(max + min)
						break LOOP
					}
				}
			}
		}
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
