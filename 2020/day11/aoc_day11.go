// https://adventofcode.com/2020/day11
//
// Read the file from stdin

package main

import (
	"bufio"
	"fmt"
	"os"
)

type matrix []string

var max_row int
var max_col int

func show(a matrix) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func count_occupied(a matrix) int {
	occupied := 0
	for _, v := range a {
		for _, v1 := range v {
			if v1 == '#' {
				occupied++
			}
		}
	}
	return occupied
}

// Follow the movement rules until we hit edge of array, or find an empty or occupied seat
func test_axis(a matrix, x int, y int, dx int, dy int) int {

	for true {
		x += dx
		y += dy
		if x < 0 || x >= max_col {
			return 0
		}
		if y < 0 || y >= max_row {
			return 0
		}
		if a[x][y] == 'L' {
			return 0
		}
		if a[x][y] == '#' {
			return 1
		}
	}
	return 0
}

// Count how many occupied seats there are around this seat along all axis
func test_whole(a matrix, x int, y int) int {
	occupied := 0

	// Test each axis
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			occupied += test_axis(a, x, y, dx, dy)
		}
	}

	return occupied
}

// Count how many occupied seats there are immediately around this seat
func test_immediate(a matrix, x int, y int) int {
	occupied := 0
	for r := x - 1; r <= x+1; r++ {
		for c := y - 1; c <= y+1; c++ {
			if r < 0 || r >= max_row {
				continue
			}
			if c < 0 || c >= max_col {
				continue
			}
			if r == x && c == y {
				continue
			}
			if a[r][c] == '#' {
				occupied++
			}
		}
	}
	return occupied
}

func iterate(test func(matrix, int, int) int, a matrix, occupied_count int) (matrix, bool) {
	var next matrix
	var s string
	changed := false

	for k, v := range a {
		s = ""
		for k1, v1 := range v {
			// Count number occupied around us
			occupied := test(a, k, k1)

			new_state := v1

			if v1 == 'L' && occupied == 0 { // empty seat
				new_state = '#'
			}
			if v1 == '#' && occupied >= occupied_count { // occupied seat
				new_state = 'L'
			}
			s += string(new_state)
			if new_state != v1 {
				changed = true
			}
		}
		next = append(next, s)
	}

	return next, changed
}

func main() {

	var array matrix

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		array = append(array, line)
		max_col = len(line)
	}

	max_row = len(array)
	fmt.Println(max_row, max_col)

	original := array

	count := 0
	changed := true
	for true {
		// fmt.Println("Iteration #", count)
		// show(array)

		array, changed = iterate(test_immediate, array, 4)

		if !changed {
			break
		}

		count++
	}
	fmt.Println("pt1 occupied seats =", count_occupied(array))

	array = original
	count = 0
	changed = true
	for true {
		// fmt.Println("Iteration #", count)
		// show(array)

		array, changed = iterate(test_whole, array, 5)

		if !changed {
			break
		}

		count++
	}
	fmt.Println("occupied seats =", count_occupied(array))

}
