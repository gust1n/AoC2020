package main

import "fmt"

func day2(input []int) {
LOOP:
	for _, n1 := range input {
		// try with each other inpt
		for _, n2 := range input {
			if n1 == n2 { // skip self
				continue
			}

			for _, n3 := range input {
				if n1 == n2 || n2 == n3 || n1 == n3 { // skip self
					continue
				}

				if n1+n2+n3 == 2020 {
					fmt.Printf("found %d, %d and %d, with the product %d\n", n1, n2, n3, n1*n2*n3)
					break LOOP
				}
			}
		}
	}
}
