package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type seat struct {
	row, col int64
}

func main() {
	lines := readStringsFromFile("./input.txt")
	var seats []seat

	for _, line := range lines {
		rowStr := line[0:7]
		colStr := line[7:10]

		rowBinary := strings.Replace(strings.Replace(rowStr, "F", "0", -1), "B", "1", -1)
		rowDecimal, err := strconv.ParseInt(rowBinary, 2, 64)
		check(err)

		colBinary := strings.Replace(strings.Replace(colStr, "L", "0", -1), "R", "1", -1)
		colDecimal, err := strconv.ParseInt(colBinary, 2, 64)
		check(err)

		seats = append(seats, seat{row: rowDecimal, col: colDecimal})
	}

	func() {
		fmt.Println("----part 1----")

		var max int64
		for _, seat := range seats {
			seatID := (seat.row * 8) + seat.col
			if seatID > max {
				max = seatID
			}
		}

		fmt.Println(max)
	}()

	func() {
		fmt.Println("----part 2----")

		var seatIDs []int
		for _, seat := range seats {
			seatID := (int(seat.row) * 8) + int(seat.col)
			seatIDs = append(seatIDs, seatID)
		}
		sort.Ints(seatIDs)

		i := seatIDs[0]
		for _, seatID := range seatIDs {
			if seatID != i {
				fmt.Println(i)
				break
			}
			i++
		}
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
