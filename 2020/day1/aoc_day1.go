// https://adventofcode.com/2020/day/1
//
// Find pairs that sum to 2020, then work out their product
//
// Read all values - put them in map
// Iterate through map. We can work out the difference between each value at 2020 - see if that is in the map
// If it is, then we have our two numbers. Work out their product
//
// Using a map provides us an efficient hash search of the number we are looking for - we could also just loop
// through the list of numbers we had to see if we've found it. Could be done quicker if we sorted it
//
// Read the file from stdin

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const sum int = 2020

// We could have a map of int to bool; this approach is more like a set where we have a map of int to nothing
type void struct{}

var member void
var vals = map[int]void{}

// Search for two numbers that sum to our required value
func searchForSum(s int) int {
	for k, _ := range vals {
		req := s - k

		// Test to see if req is in our map - this works because we know none of our values is 0
		_, exists := vals[req]
		if exists {
			return k
		}
	}
	return 0
}

func main() {

	// Parse the input stream to create a map of integers
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err == nil {
			vals[val] = member
		}
	}

	// Find the pair that sum to <sum>
	second := searchForSum(sum)
	fmt.Println("Found our pair: ", sum-second, second)
	fmt.Println("Sum is ", (sum-second)*second)

	// Try each value in the map to find the three values that make our required sum

	for k, _ := range vals {
		req := sum - k

		second := searchForSum(req)

		if second != 0 {
			third := req - second
			fmt.Println("Found our three numbers: ", k, second, third)
			fmt.Println("Sum is ", k*second*third)
			break
		}
	}

}
