package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := readStringsFromFile("./input.txt")
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	func() {
		fmt.Println("----part 1----")

		var currentIdx int
		var numValid int
		var passports [][]string
		for i, line := range lines {
			if len(line) == 0 { // new passport
				var passport []string
				for j := currentIdx; j < i; j++ { // each line for this passport
					passport = append(passport, strings.Fields(lines[j])...)
				}
				currentIdx = i + 1
				passports = append(passports, passport)
			}
			if i+1 == len(lines) { // last line
				var passport []string
				for j := currentIdx; j < len(lines); j++ { // each line for this passport
					passport = append(passport, strings.Fields(lines[j])...)
				}
				passports = append(passports, passport)
			}
		}
		for _, passport := range passports {
			if hasRequiredFields(passport, requiredFields) {
				numValid++
			}
		}

		fmt.Println(numValid)
	}()

	func() {
		fmt.Println("----part 2----")

		var currentIdx int
		var numValid int
		var passports [][]string
		for i, line := range lines {
			if len(line) == 0 { // new passport
				var passport []string
				for j := currentIdx; j < i; j++ { // each line for this passport
					passport = append(passport, strings.Fields(lines[j])...)
				}
				currentIdx = i + 1
				passports = append(passports, passport)
			}
			if i+1 == len(lines) { // last line
				var passport []string
				for j := currentIdx; j < len(lines); j++ { // each line for this passport
					passport = append(passport, strings.Fields(lines[j])...)
				}
				passports = append(passports, passport)
			}
		}

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

// rules:

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
//     If cm, the number must be at least 150 and at most 193.
//     If in, the number must be at least 59 and at most 76.
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.

var hclReg = regexp.MustCompile(`#[a-f0-9]{6}`)
var pidReg = regexp.MustCompile(`[0-9]{9}`)

func hasValidValues(passport []string) bool {
	keys := map[string]string{}
	for _, pair := range passport {
		kv := strings.Split(pair, ":")
		keys[kv[0]] = kv[1]
	}

	// byr
	byr := parseInt(keys["byr"])
	if byr < 1920 || byr > 2002 {
		return false
	}

	// iyr
	iyr := parseInt(keys["iyr"])
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	// eyr
	eyr := parseInt(keys["eyr"])
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	// hgt
	hgt := keys["hgt"]
	unit := hgt[max(0, len(hgt)-2):len(hgt)]
	if unit == "cm" {
		value := parseInt(hgt[0 : len(hgt)-2])
		if value < 150 || value > 193 {
			return false
		}
	} else if unit == "in" {
		value := parseInt(hgt[0 : len(hgt)-2])
		if value < 59 || value > 76 {
			return false
		}
	} else {
		return false
	}

	// hcl
	hcl := keys["hcl"]
	if !hclReg.Match([]byte(hcl)) {
		return false
	}

	// ecl
	ecl := keys["ecl"]
	if len(ecl) != 3 || !strings.Contains("amb blu brn gry grn hzl oth", ecl) {
		return false
	}

	// pid
	pid := keys["pid"]
	if !pidReg.Match([]byte(pid)) {
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
