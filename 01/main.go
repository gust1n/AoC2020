package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines := readIntLinesFromFile("./input.txt")

	func() {
		fmt.Println("----part 1----")

	LOOP:
		for _, n1 := range lines {
			for _, n2 := range lines {
				if n1+n2 == 2020 {
					fmt.Printf("found %d and %d, with the product %d\n", n1, n2, n1*n2)
					break LOOP
				}
			}
		}

	}()

	func() {
		fmt.Println("----part 2----")

	LOOP:
		for _, n1 := range lines {
			for _, n2 := range lines {
				for _, n3 := range lines {
					if n1+n2+n3 == 2020 {
						fmt.Printf("found %d, %d and %d, with the product %d\n", n1, n2, n3, n1*n2*n3)
						break LOOP
					}
				}
			}
		}

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

func readIntLinesFromFile(path string) []int {
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
