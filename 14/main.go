package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const maskPrefix = "mask = "

var memRegexp = regexp.MustCompile(`^mem\[(\d+)\]\s\=\s(\d+)$`)

func main() {
	lines := readStringsFromFile("./input.txt")

	func() {
		fmt.Println("----part 1----")

		mem := map[int]int64{}

		var maskWithZeroes, maskWithOnes int64
		for _, line := range lines {
			if strings.HasPrefix(line, maskPrefix) {
				mask := strings.TrimPrefix(line, maskPrefix)
				maskWithZeroes = binToInt64(strings.Replace(mask, "X", "0", -1))
				maskWithOnes = binToInt64(strings.Replace(mask, "X", "1", -1))
			} else {
				matches := memRegexp.FindStringSubmatch(line)
				addr, decimalValue := parseInt(matches[1]), int64(parseInt(matches[2]))

				mem[addr] = (decimalValue | maskWithZeroes) & maskWithOnes
			}
		}

		var sum int64
		for _, val := range mem {
			sum += val
		}
		fmt.Println(sum)
	}()

	func() {
		fmt.Println("----part 2----")

		mem := map[int64]int64{}

		var mask string
		for _, line := range lines {
			if strings.HasPrefix(line, maskPrefix) {
				mask = strings.TrimPrefix(line, maskPrefix)
			} else {
				matches := memRegexp.FindStringSubmatch(line)
				addr, decimalValue := parseInt(matches[1]), int64(parseInt(matches[2]))
				addrString := fmt.Sprintf("%036b", addr)

				// apply mask manually
				var b strings.Builder
				for i := 0; i < 36; i++ {
					val := addrString[i]
					if mask[i] != '0' {
						val = mask[i]
					}
					fmt.Fprintf(&b, "%s", string(val))
				}
				res := b.String()

				// find index of all X-s
				var xIndexes []int
				for i, c := range res {
					if c == 'X' {
						xIndexes = append(xIndexes, i)
					}
				}

				// generate power set of all combinations of ones
				for _, oneIndexes := range powerSet(xIndexes) {
					// current result (but X replaced with 0)
					var chars []rune
					for _, c := range res {
						if c == 'X' {
							chars = append(chars, '0')
						} else {
							chars = append(chars, c)
						}
					}

					// replace current combination of 1-indexes with 1
					for _, oneIndex := range oneIndexes {
						chars[oneIndex] = '1'
					}

					decimalAddr, err := strconv.ParseInt(string(chars), 2, 64)
					check(err)
					mem[decimalAddr] = decimalValue
				}
			}
		}

		var sum int64
		for _, val := range mem {
			sum += val
		}
		fmt.Println(sum)
	}()
}

func powerSet(original []int) [][]int {
	powerSetSize := int(math.Pow(2, float64(len(original))))
	result := make([][]int, 0, powerSetSize)

	var index int
	for index < powerSetSize {
		var subSet []int

		for j, elem := range original {
			if index&(1<<uint(j)) > 0 {
				subSet = append(subSet, elem)
			}
		}
		result = append(result, subSet)
		index++
	}
	return result
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func binToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 2, 64)
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
