package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var instructionRegexp = regexp.MustCompile(`^(jmp|acc|nop)\s(\+|\-)(\d+)$`)

type instruction struct {
	kind  string
	value int
}

func runProgram(program map[int]instruction) (int, bool) {
	pc := 0
	globalAcc := 0
	taken := map[int]bool{}

	for {
		instruction := program[pc]

		// if end at program
		if pc == len(program) {
			break
		}

		// if revisiting
		if _, ok := taken[pc]; ok {
			break
		}
		taken[pc] = true

		switch instruction.kind {
		case "nop":
		case "jmp":
			pc += instruction.value
			continue
		case "acc":
			globalAcc += instruction.value
		}

		pc++
	}

	return globalAcc, pc == len(program)
}

func main() {
	lines := readStringsFromFile("./input.txt")

	func() {
		fmt.Println("----part 1----")

		program := makeProgram(lines)
		acc, _ := runProgram(program)

		fmt.Println(acc)
	}()

	func() {
		fmt.Println("----part 2----")

		program := makeProgram(lines)

		// build map of instructions (index -> new instruction) to try changing
		instructionsToChange := map[int]string{}
		for i, instruction := range program {
			switch instruction.kind {
			case "acc": // acc are kept
			case "nop":
				instructionsToChange[i] = "jmp"
			case "jmp":
				instructionsToChange[i] = "nop"
			}
		}

		for index, newInstruction := range instructionsToChange {
			program := makeProgram(lines)
			program[index] = instruction{
				kind:  newInstruction,
				value: program[index].value,
			}

			res, completed := runProgram(program)
			if completed {
				fmt.Println(res)
				break
			}
		}
	}()
}

func makeProgram(lines []string) map[int]instruction {
	program := map[int]instruction{}

	for i, line := range lines {
		matches := instructionRegexp.FindStringSubmatch(line)
		kind := matches[1]
		sign := matches[2]
		value := parseInt(matches[3])
		if sign == "-" {
			value = -value
		}
		program[i] = instruction{
			kind:  kind,
			value: value,
		}

	}

	return program
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
