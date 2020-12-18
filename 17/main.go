package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type grid map[vector]struct{}

func (g grid) min() vector {
	mv := vector{
		x: math.MaxInt64,
		y: math.MaxInt64,
		z: math.MaxInt64,
		w: math.MaxInt64,
	}

	for v := range g {
		mv.x = min(mv.x, v.x)
		mv.y = min(mv.y, v.y)
		mv.z = min(mv.z, v.z)
		mv.w = min(mv.w, v.w)
	}

	return mv
}

func (g grid) max() vector {
	mv := vector{
		x: math.MinInt64,
		y: math.MinInt64,
		z: math.MinInt64,
		w: math.MinInt64,
	}

	for v := range g {
		mv.x = max(mv.x, v.x)
		mv.y = max(mv.y, v.y)
		mv.z = max(mv.z, v.z)
		mv.w = max(mv.w, v.w)
	}

	return mv
}

func newGrid() grid {
	return map[vector]struct{}{}
}

func main() {
	lines := readStringsFromFile("./input.txt")

	// Parse out active vectors from input
	initialActiveVectors := newGrid()
	for row, line := range lines {
		for col, char := range line {
			if char == '#' {
				initialActiveVectors[vector{y: row, x: col}] = struct{}{}
			}
		}
	}

	func() {
		fmt.Println("----part 1----")

		processGrid := func(g grid) grid {
			ng := newGrid()

			minVector := g.min().add(vector{x: -1, y: -1, z: -1})
			maxVector := g.max().add(vector{x: 1, y: 1, z: 1})

			for z := minVector.z; z <= maxVector.z; z++ {
				for y := minVector.y; y <= maxVector.y; y++ {
					for x := minVector.x; x <= maxVector.x; x++ {
						v := vector{x: x, y: y, z: z}
						_, isActive := g[v]

						var neighboursActive int
						for _, n := range v.neighbours3d() {
							if _, ok := g[n]; ok {
								neighboursActive++
							}
						}

						if isActive && ((neighboursActive == 2) || (neighboursActive == 3)) {
							ng[v] = struct{}{}
						} else if !isActive && (neighboursActive == 3) {
							ng[v] = struct{}{}
						}
					}
				}
			}

			return ng
		}

		activeVectors := initialActiveVectors
		for i := 1; i <= 6; i++ {
			activeVectors = processGrid(activeVectors)
		}

		fmt.Println(len(activeVectors))
	}()

	func() {
		fmt.Println("----part 2----")

		processGrid := func(g grid) grid {
			ng := newGrid()

			minVector := g.min().add(vector{x: -1, y: -1, z: -1, w: -1})
			maxVector := g.max().add(vector{x: 1, y: 1, z: 1, w: 1})

			for w := minVector.w; w <= maxVector.w; w++ {
				for z := minVector.z; z <= maxVector.z; z++ {
					for y := minVector.y; y <= maxVector.y; y++ {
						for x := minVector.x; x <= maxVector.x; x++ {
							v := vector{x: x, y: y, z: z, w: w}
							_, isActive := g[v]

							var neighboursActive int
							for _, n := range v.neighbours4d() {
								if _, ok := g[n]; ok {
									neighboursActive++
								}
							}

							if isActive && ((neighboursActive == 2) || (neighboursActive == 3)) {
								ng[v] = struct{}{}
							} else if !isActive && (neighboursActive == 3) {
								ng[v] = struct{}{}
							}
						}
					}
				}
			}

			return ng
		}

		activeVectors := initialActiveVectors
		for i := 1; i <= 6; i++ {
			activeVectors = processGrid(activeVectors)
		}

		fmt.Println(len(activeVectors))
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

type vector struct {
	x, y, z, w int
}

func (v vector) add(v2 vector) vector {
	return vector{
		x: v.x + v2.x,
		y: v.y + v2.y,
		z: v.z + v2.z,
		w: v.w + v2.w,
	}
}

func (v vector) equal(v2 vector) bool {
	return (v.x == v2.x) && (v.y == v2.y) && (v.z == v2.z) && (v.w == v2.w)
}

func (v vector) neighbours3d() []vector {
	var n []vector
	for z := v.z - 1; z <= (v.z + 1); z++ {
		for y := v.y - 1; y <= (v.y + 1); y++ {
			for x := v.x - 1; x <= (v.x + 1); x++ {
				v2 := vector{x: x, y: y, z: z}
				if !v2.equal(v) { // skip self
					n = append(n, v2)
				}
			}
		}
	}

	return n
}

func (v vector) neighbours4d() []vector {
	var n []vector
	for w := v.w - 1; w <= (v.w + 1); w++ {
		for z := v.z - 1; z <= (v.z + 1); z++ {
			for y := v.y - 1; y <= (v.y + 1); y++ {
				for x := v.x - 1; x <= (v.x + 1); x++ {
					v2 := vector{x: x, y: y, z: z, w: w}
					if !v2.equal(v) { // skip self
						n = append(n, v2)
					}
				}
			}
		}
	}

	return n
}

func (v vector) sub(v2 vector) vector {
	return vector{
		x: v.x - v2.x,
		y: v.y - v2.y,
		z: v.z - v2.z,
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(list ...int) int {
	sort.Ints(list)
	return list[0]
}

func max(list ...int) int {
	sort.Ints(list)
	return list[len(list)-1]
}
