// https://adventofcode.com/2020/day10
//
// Read the file from stdin

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var list_size int

func main() {

	var list []int

	// We start at 0
	list = append(list, 0)

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		v, _ := strconv.Atoi(line)
		list = append(list, v)
	}
	fmt.Println(list)

	// A simple sort works because the differences are all <=3. Note there are other sorting
	// options that would work
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	// And we end 3 more than the final value
	list = append(list, list[len(list)-1]+3)

	fmt.Println(list)

	var diff [4]int
	var runs [2]int

	len_of_run := 0
	for k, v := range list {
		if k == 0 {
			continue
		}
		d := v - list[k-1]
		if d == 1 {
			len_of_run++
		}
		if d == 3 && len_of_run != 0 {
			fmt.Println("Found a run of", len_of_run, "1's")

			// TODO: This count rules works for the test set but not the full data set... need to think more!
			if len_of_run > 3 {
				runs[1]++
			} else {
				runs[0]++
			}
			len_of_run = 0
		}
		fmt.Println(k, v, d)
		if d > 3 {
			fmt.Println("Bad list sorting", v, "-", list[k-1], ">3")
		} else {
			diff[d]++
		}
	}
	fmt.Println(diff)
	// Part 1
	fmt.Println("# 1-jolt =", diff[1], "# 3-jolt", diff[3])
	fmt.Println("Product = ", diff[1]*diff[3])

	// Part 2 - this works for the test set, but fails for the data set
	fmt.Println("runs", runs)
	fmt.Printf("calc %.0f\n", math.Pow(2, float64(runs[0]))*math.Pow(7, float64(runs[1])))
}
