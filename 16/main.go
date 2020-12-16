package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var ruleRegexp = regexp.MustCompile(`^(.+):\s(\d+)-(\d+)\sor\s(\d+)-(\d+)$`)
var ticketRegecp = regexp.MustCompile(`^\d+,`)

type rule struct {
	name  string
	valid map[int]struct{}
}

func main() {
	lines := readStringsFromFile("./input.txt")

	var rules []rule
	var tickets [][]int

	for _, line := range lines {
		if matches := ruleRegexp.FindStringSubmatch(line); matches != nil {
			r := rule{name: matches[1], valid: map[int]struct{}{}}

			// fill with all valid numbers for easy lookup
			for i := parseInt(matches[2]); i <= parseInt(matches[3]); i++ {
				r.valid[i] = struct{}{}
			}
			for i := parseInt(matches[4]); i <= parseInt(matches[5]); i++ {
				r.valid[i] = struct{}{}
			}

			rules = append(rules, r)
		} else if matches := ticketRegecp.MatchString(line); matches {
			numStrings := strings.Split(line, ",")
			var nums []int
			for _, s := range numStrings {
				nums = append(nums, parseInt(s))
			}
			tickets = append(tickets, nums)
		}
	}

	allValidNums := map[int]struct{}{}

	for _, r := range rules {
		for num := range r.valid {
			allValidNums[num] = struct{}{}
		}
	}

	func() {
		fmt.Println("----part 1----")

		var errorRate int
		for _, t := range tickets[1:] { // my ticket is first, skip it
			for _, num := range t {
				if _, ok := allValidNums[num]; !ok {
					errorRate += num
				}
			}
		}

		fmt.Println(errorRate)
	}()

	func() {
		fmt.Println("----part 2----")

		// First, filter out the valid tickets
		var validTickets [][]int

		for _, t := range tickets {
			valid := true
			for _, num := range t {
				if _, ok := allValidNums[num]; !ok {
					valid = false
					break
				}
			}
			if valid {
				validTickets = append(validTickets, t)
			}
		}

		// Second, put each field in columns
		cols := make([][]int, len(tickets[0]))
		for row, t := range validTickets {
			for col, num := range t {
				if cols[col] == nil {
					cols[col] = make([]int, len(validTickets))
				}
				cols[col][row] = num
			}
		}

		// Third, check which cols are valid according to each rule
		ruleCols := map[string][]int{}
		for colIdx, col := range cols {
		LOOP:
			for _, rule := range rules {
				for _, num := range col {
					if _, ok := rule.valid[num]; !ok {
						continue LOOP
					}
				}
				ruleCols[rule.name] = append(ruleCols[rule.name], colIdx)
			}
		}

		// Fourth, find columns that are only valid for one rule and remove
		// until found all rules
		foundCols := map[string]int{}
		for len(foundCols) < len(ruleCols) {
			for ruleName, cols := range ruleCols {
				if len(cols) == 1 {
					foundCol := cols[0]
					foundCols[ruleName] = foundCol

					// remove from all rules
					for rule, cols := range ruleCols {
						ruleCols[rule] = filterInts(cols, func(i int) bool {
							return i != foundCol
						})
					}
					break
				}
			}
		}

		prod := 1
		for rule, col := range foundCols {
			if strings.HasPrefix(rule, "departure") {
				prod = prod * validTickets[0][col]
			}
		}
		fmt.Println(prod)
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
