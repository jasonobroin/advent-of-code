// https://adventofcode.com/2020/day9
//
// Read the file from stdin

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var list_size int

func init() {
	const (
		defaultListSize = 25
		usage           = "Number of entries in preamble"
	)

	flag.IntVar(&list_size, "p", defaultListSize, usage)

}

func testsum(val int64, list []int64) bool {
	for k, v := range list {
		for k1, v1 := range list {
			if k != k1 { // Don't sum ourselves
				if v+v1 == val {
					return true
				}
			}
		}
	}
	return false
}

func main() {

	flag.Parse()

	var list []int64
	preload := list_size

	list_start := 0
	list_end := 0

	var missing int64

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		v, _ := strconv.ParseInt(line, 10, 64)
		list = append(list, v)

		// Preload a certain number before we do our searches
		if preload > 0 {
			preload--
		} else {
			// Test if v meets our criteria (its a sum of previous values in the array
			res := testsum(v, list[list_start:list_end])
			if !res {
				fmt.Println(v, "not in list")
				missing = v
			}
		}

		list_end++
		if list_end-list_start > list_size {
			list_start++
		}

	}

	// Find run of values that sum to <missing>
	for k, v := range list {
		if v >= missing {
			// Can't make the sum with 2+ numbers
			continue
		}
		var total int64 = 0
		i := k
		for total < missing {
			total += list[i]
			if total == missing {
				// Find the smallest and largest in the range that makes our sum
				var s int64 = list[k]
				var l int64 = list[k]
				for _, b := range list[k:i] {
					if b < s {
						s = b
					}
					if b > l {
						l = b
					}
				}

				fmt.Println(list[k], "...", list[i], "=", missing, "sum=", s+l)
			}
			i++

		}
	}

}
