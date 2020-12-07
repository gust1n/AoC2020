package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var lineRegexp = regexp.MustCompile(`^(\w+\s\w+)\sbags\scontain\s(.*)`)
var includesRegexp = regexp.MustCompile(`^(\d)\s(\w+\s\w+)\sbag[s]?`)

func containsColor(bagTypes map[string]map[string]int, outerBagColor, targetColor string) bool {
	if _, ok := bagTypes[outerBagColor][targetColor]; ok {
		return true
	}

	for c := range bagTypes[outerBagColor] {
		if containsColor(bagTypes, c, targetColor) {
			return true
		}
	}

	return false
}

func checkDepth(bagTypes map[string]map[string]int, depth int, color string) int {
	for c, num := range bagTypes[color] {
		depth += num * checkDepth(bagTypes, 1, c)
	}

	return depth
}

func main() {
	lines := readStringsFromFile("./input.txt")

	// outer bag color -> inner bag color -> number
	bagTypes := map[string]map[string]int{}

	for _, line := range lines {
		matches := lineRegexp.FindStringSubmatch(line)

		outerColor := matches[1]
		rest := matches[2]

		if rest == "no other bags." {
			bagTypes[outerColor] = nil
			continue
		}

		innerBags := map[string]int{}

		for _, includeStr := range strings.Split(rest, ",") {
			innerMatches := includesRegexp.FindStringSubmatch(strings.TrimSpace(includeStr))
			num := parseInt(innerMatches[1])
			color := innerMatches[2]

			innerBags[color] = num
		}

		bagTypes[outerColor] = innerBags
	}

	func() {
		fmt.Println("----part 1----")

		goldCount := 0

		for outerBagColor := range bagTypes {
			if containsColor(bagTypes, outerBagColor, "shiny gold") {
				goldCount++
			}
		}

		fmt.Println(goldCount)
	}()

	func() {
		fmt.Println("----part 2----")

		var totalDepth int
		for innerBagColor, num := range bagTypes["shiny gold"] {
			totalDepth += num * checkDepth(bagTypes, 1, innerBagColor)
		}

		fmt.Println(totalDepth)
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
