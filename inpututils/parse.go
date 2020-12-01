package inpututils

import (
	"bufio"
	"bytes"
	"strconv"
)

// ParseInts parses an input list into a Go type int slice.
func ParseInts(input []byte) ([]int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(input))
	var res []int
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) == 0 {
			continue
		}

		i, err := strconv.Atoi(t)
		if err != nil {
			return nil, err
		}

		res = append(res, i)
	}

	return res, nil
}
