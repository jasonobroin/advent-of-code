// https://adventofcode.com/2020/day/3
//
// Find our how many trees we encounter traversing our array
//
// Our traversal strategy is right 3, down 1
//
// Our array wraps continuously as we hit the right edge
//
// We have an array of empty and filled spot. We want to count
// the number of filled spots (which represent trees) as we move
// through the array. Each line is the same size
//
// Read the file from stdin

package main

import (
	"bufio"
	"fmt"
	"os"
)

var linelength = 0
var maxrows = 0
var matrix = make([][]bool, 0)

func navigate_slope(x_move int, y_move int) int {
	// Navigate array

	// current location. Start at top left

	x := 0
	y := 0

	trees := 0

	for y < maxrows {
		if matrix[y][x] == true {
			trees += 1
		}
		x += x_move
		y += y_move

		if x >= linelength {
			x -= linelength
		}
	}

	fmt.Println("x move ", x_move, "y move", y_move, "Hit ", trees, "trees")

	return trees
}

func main() {

	// Parse the input stream to create a map of bools
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if linelength == 0 {
			linelength = len(line)
			fmt.Println("line length", linelength)
		}

		row := make([]bool, linelength)

		for i, c := range line {
			if c == '#' {
				row[i] = true // a tree
			} else {
				row[i] = false
			}
		}

		matrix = append(matrix, row)

	}

	maxrows = len(matrix)
	fmt.Println("matrix rows", maxrows)

	run1 := navigate_slope(1, 1)
	run2 := navigate_slope(5, 1)
	run3 := navigate_slope(7, 1)
	run4 := navigate_slope(1, 2)
	run5 := navigate_slope(3, 1)
	fmt.Println(run1 * run2 * run3 * run4 * run5)

}
