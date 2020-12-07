// https://adventofcode.com/2020/day6/
//
// Read the file from stdin

package main

import (
	"bufio"
	"fmt"
	"os"
)

const maxchar int = 26

func reset(a *[maxchar]int) {
	for k, _ := range a {
		a[k] = 0
	}
}

func main() {

	part1_total := 0
	part2_total := 0
	var answers [maxchar]int
	part1_count := 0
	people := 0

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		people++

		for _, v := range line {
			// Test before we set so we can get our count on one loop through
			if answers[v-'a'] == 0 {
				part1_count++
			}
			answers[v-'a']++
		}

		if len(line) == 0 {
			people--
			// How many questions did one person answer yes
			part1_total += part1_count

			// How many questions did everyone answer yes

			for _, v := range answers {
				if v == people {
					part2_total++
				}
			}

			reset(&answers)
			part1_count = 0
			people = 0

		}
	}

	fmt.Println("part1 sum of counts", part1_total)
	fmt.Println("part2 total", part2_total)
}
