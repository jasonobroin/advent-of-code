// https://adventofcode.com/2020/day2/
//
// Determine how many passwords are valid in file
//
// Each line has a format
//
// 1-3 a: abcde
//
// <min>-<max> <char>: <password>
//
// The letter <char> must appear >= min, and <= max times in the password
//
// We need to look at each line, count the number of occurences of <char> and see if it matches the policy
//
// Read the file from stdin

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	total_policy1 := 0
	total_policy2 := 0

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Split string into component pieces
		first := strings.Split(line, ":")
		password := first[1]
		second := strings.Split(first[0], " ")
		searchchar := second[1]
		third := strings.Split(second[0], "-")
		min, _ := strconv.Atoi(third[0])
		max, _ := strconv.Atoi(third[1])

		// Count how many instances of searchchar in our password
		count := strings.Count(password, searchchar)

		// See if we match the policy 1 rule

		if count >= min && count <= max {
			total_policy1 += 1
			//			fmt.Println("line: ", line, " = ", min, ",", max, ",", searchchar, ",", password, ",", count)
		}

		// See if we match the policy 2 rule

		firstpos := min
		secondpos := max
		numfound := 0
		if string(password[firstpos]) == searchchar {
			numfound += 1
		}
		if string(password[secondpos]) == searchchar {
			numfound += 1
		}
		if numfound == 1 {
			total_policy2 += 1
			//			fmt.Println("line: ", line, " = ", min, ",", max, ",", searchchar, ",", password)
		}

	}
	fmt.Println("Found ", total_policy1, " correct matches (policy 1)")
	fmt.Println("Found ", total_policy2, " correct matches (policy 2)")
}
