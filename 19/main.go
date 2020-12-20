package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var ruleRegexp = regexp.MustCompile(`^(\d+):\s(\d.*)`)
var messageRegexp = regexp.MustCompile(`^[ab]+`)
var numRegexp = regexp.MustCompile(`(\d+)`)

func processRuleRecursion(aIdx, bIdx int, ruleIdx int, ruleStrings map[int]string) string {
	ruleStr := ruleStrings[ruleIdx]
	// fmt.Printf("processing rule %d, %s\n", ruleIdx, ruleStr)

	// find sub rules
	matches := numRegexp.FindAllString(ruleStr, -1)
	for _, subRuleString := range matches {
		subRuleNum := parseInt(subRuleString)
		// fmt.Println("processing subrule", subRuleNum)
		if subRuleNum == aIdx {
			ruleStr = strings.Replace(ruleStr, subRuleString, fmt.Sprintf("([a])"), 1)
		} else if subRuleNum == bIdx {
			ruleStr = strings.Replace(ruleStr, subRuleString, fmt.Sprintf("([b])"), 1)
		} else {
			newRule := processRuleRecursion(aIdx, bIdx, subRuleNum, ruleStrings)
			ruleStr = strings.Replace(ruleStr, subRuleString, fmt.Sprintf("(%s)", newRule), 1)
		}
	}

	return strings.Replace(ruleStr, " ", "", -1)
}

func main() {
	lines := readStringsFromFile("./input.txt")

	ruleStrings := map[int]string{}
	var messages []string

	for _, line := range lines {
		if matches := ruleRegexp.FindStringSubmatch(line); matches != nil {
			ruleNum := parseInt(matches[1])
			ruleStr := matches[2]

			ruleStrings[ruleNum] = ruleStr
		} else if match := messageRegexp.FindString(line); match != "" {
			messages = append(messages, line)
		}
	}

	func() {
		fmt.Println("----part 1----")

		rule0RegexpStr := processRuleRecursion(72, 58, 0, ruleStrings)
		rule0regexp := regexp.MustCompile(fmt.Sprintf("^%s$", rule0RegexpStr))
		var numFullMatches int
		for _, msg := range messages {
			if rule0regexp.MatchString(msg) {
				numFullMatches++
			}
		}
		fmt.Println(numFullMatches)
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
