package main

import "fmt"

func day1(input []int) {
LOOP:
	for _, n1 := range input {
		for _, n2 := range input {
			if n1+n2 == 2020 {
				fmt.Printf("found %d and %d, with the product %d\n", n1, n2, n1*n2)
				break LOOP
			}
		}
	}
}
