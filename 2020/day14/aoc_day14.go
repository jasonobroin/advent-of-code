// https://adventofcode.com/2020/day13
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

var bits uint = 36

func masked_val(val uint64, m string) uint64 {
	for k, v := range m {
		if v == '1' { // set bit
			val |= 1 << (bits - uint(k) - 1)
		}
		if v == '0' { // clear bit
			val &^= 1 << (bits - uint(k) - 1)
		}
	}
	return val
}

func masked_loc(loc uint64, m []byte, a []uint64) []uint64 {
	for k, v := range m {
		if v == '1' { // set bit
			loc |= 1 << (bits - uint(k) - 1)
		}
		if v == 'X' { // split
			loc |= 1 << (bits - uint(k) - 1)
			s := m
			s[k] = '1'
			if uint(k) == bits-1 {
				a = append(a, loc)
			} else {
				a = masked_loc(loc, s, a)
			}

			loc &^= 1 << (bits - uint(k) - 1)
			s[k] = '0'
			if uint(k) == bits-1 {
				a = append(a, loc)
			} else {
				a = masked_loc(loc, s, a)
			}
			s[k] = 'X' // Restore X as we don't have a copy of the array, but a slice
			return a
		}
	}
	a = append(a, loc)
	return a
}

func main() {

	mask := ""

	var memory_pt1 = map[int]uint64{}
	var memory_pt2 = map[uint64]uint64{}

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if line[0:4] == "mask" {
			mask = line[7:]
		} else {
			first := strings.Split(line, "=")
			value, _ := strconv.Atoi(strings.TrimSpace(first[1]))
			second := strings.Split(first[0], "[")
			third := strings.Split(second[1], "]")
			loc, _ := strconv.Atoi(strings.TrimSpace(third[0]))

			//pt1
			mv := masked_val(uint64(value), mask)
			memory_pt1[loc] = mv

			// pt2
			var ml []uint64
			ml = masked_loc(uint64(loc), []byte(mask), ml)
			for _, v := range ml {
				memory_pt2[uint64(v)] = uint64(value)
			}

		}
	}

	var total_pt1 uint64 = 0
	for _, v := range memory_pt1 {
		total_pt1 += v
	}

	var total_pt2 uint64 = 0
	for _, v := range memory_pt2 {
		total_pt2 += v
	}

	//	fmt.Println(memory)
	fmt.Println("Memory total pt1", total_pt1)
	fmt.Println("Memory total pt2", total_pt2)
}
