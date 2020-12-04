package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport map[string]string

func main() {
	lines := readStringsFromFile("./input.txt")
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	var entryStartIdx int
	var passports [][]string

	for i, line := range lines {
		if len(line) == 0 { // new passport
			var passport []string
			for j := entryStartIdx; j < i; j++ { // each line for this passport
				passport = append(passport, strings.Fields(lines[j])...)
			}
			entryStartIdx = i + 1
			passports = append(passports, passport)
		}
		if i+1 == len(lines) { // last line
			var passport []string
			for j := entryStartIdx; j < len(lines); j++ { // each line for this passport
				passport = append(passport, strings.Fields(lines[j])...)
			}
			passports = append(passports, passport)
		}
	}

	func() {
		fmt.Println("----part 1----")

		var numValid int

		for _, passport := range passports {
			if hasRequiredFields(passport, requiredFields) {
				numValid++
			}
		}

		fmt.Println(numValid)
	}()

	func() {
		fmt.Println("----part 2----")

		var numValid int

		for _, passport := range passports {
			if hasRequiredFields(passport, requiredFields) && hasValidValues(passport) {
				numValid++
			}
		}

		fmt.Println(numValid)
	}()
}

func hasRequiredFields(passport []string, requiredFields []string) bool {
	keys := map[string]string{}
	for _, pair := range passport {
		kv := strings.Split(pair, ":")
		keys[kv[0]] = kv[1]
	}

	for _, requiredField := range requiredFields {
		if _, ok := keys[requiredField]; !ok {
			return false
		}
	}

	return true

}

var hclReg = regexp.MustCompile(`#[a-f0-9]{6}`)
var pidReg = regexp.MustCompile(`[0-9]{9}`)
var yearRegexp = regexp.MustCompile(`^[0-9]{4}$`)

func yearValid(s string, min, max int) bool {
	if !yearRegexp.MatchString(s) {
		return false
	}
	i := parseInt(s)
	if i < min || i > max {
		return false
	}

	return true
}

func hasValidValues(passport []string) bool {
	keys := map[string]string{}
	for _, pair := range passport {
		kv := strings.Split(pair, ":")
		keys[kv[0]] = kv[1]
	}

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	if !yearValid(keys["byr"], 1920, 2002) {
		return false
	}

	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	if !yearValid(keys["iyr"], 2010, 2020) {
		return false
	}

	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	if !yearValid(keys["eyr"], 2020, 2030) {
		return false
	}

	// hgt (Height) - a number followed by either cm or in:
	//     If cm, the number must be at least 150 and at most 193.
	//     If in, the number must be at least 59 and at most 76.
	hgt := keys["hgt"]
	if strings.HasSuffix(hgt, "cm") {
		value := parseInt(strings.TrimSuffix(hgt, "cm"))
		if value < 150 || value > 193 {
			return false
		}
	} else if strings.HasSuffix(hgt, "in") {
		value := parseInt(strings.TrimSuffix(hgt, "in"))
		if value < 59 || value > 76 {
			return false
		}
	} else {
		return false
	}

	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	hcl := keys["hcl"]
	if !hclReg.MatchString(hcl) {
		return false
	}

	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	ecl := keys["ecl"]
	if len(ecl) != 3 || !strings.Contains("amb blu brn gry grn hzl oth", ecl) {
		return false
	}

	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	pid := keys["pid"]
	if !pidReg.MatchString(pid) {
		return false
	}

	return true
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

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
