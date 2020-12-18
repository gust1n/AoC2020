package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var parenthesisRegexp = regexp.MustCompile(`\([0-9 +*]+\)`)
var additionRegexp = regexp.MustCompile(`[0-9]+\s[+]\s[0-9]+`)
var numberRegexp = regexp.MustCompile(`^[0-9]+\s?`)
var signRegexp = regexp.MustCompile(`^[+*]\s`)

func calcLeftToRight(s string) int {
	var result int
	var sign string
	for len(s) > 0 {
		if match := numberRegexp.FindString(s); len(match) > 0 {
			i := parseInt(strings.TrimSpace(match))
			switch sign {
			case "":
				result = i
			case "*":
				result *= i
			case "+":
				result += i
			}
			s = s[len(match):]
		} else if match := signRegexp.FindString(s); len(match) > 0 {
			sign = string(match[0])
			s = s[len(match):]
		}
	}

	return result
}

func calcAdditionFirst(s string) int {
	for strings.Contains(s, "+") {
		for _, m := range additionRegexp.FindAllString(s, 1) {
			res := calcLeftToRight(m)
			s = strings.Replace(s, m, fmt.Sprintf("%d", res), -1)
		}
	}

	return calcLeftToRight(s)
}

func processParenthesis(s string, calcFn func(string) int) string {
	if matches := parenthesisRegexp.FindAllString(s, -1); matches != nil {
		for _, m := range matches {
			res := calcFn(m[1 : len(m)-1])
			s = strings.Replace(s, m, fmt.Sprintf("%d", res), -1)
		}
	}

	if strings.Contains(s, "(") {
		return processParenthesis(s, calcFn)
	}

	return s
}

func main() {
	lines := readStringsFromFile("./input.txt")

	func() {
		fmt.Println("----part 1----")

		var sum int
		for _, line := range lines {
			sum += calcLeftToRight(processParenthesis(line, calcLeftToRight))
		}
		fmt.Println(sum)
	}()

	func() {
		fmt.Println("----part 2----")

		var sum int
		for _, line := range lines {
			sum += calcAdditionFirst(processParenthesis(line, calcAdditionFirst))
		}
		fmt.Println(sum)

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
