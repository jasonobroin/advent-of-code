// https://adventofcode.com/2020/day/5
//
// Seat number
//
// Read the file from stdin

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func decode(data string, f string, s string, width uint) int {
	min := 0
	max := (1 << width) - 1

	for _, l := range data {
		if string(l) == f {
			max = min + (max-min)>>1
		}
		if string(l) == s {
			min = min + ((max - min) >> 1) + 1
		}
	}

	return min
}

func main() {

	max_seat := 0
	var seats []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		row := line[0:7]
		seat := line[7:]

		r := decode(row, "F", "B", 7)
		s := decode(seat, "L", "R", 3)

		seat_id := r*8 + s

		if seat_id > max_seat {
			max_seat = seat_id
		}

		// Add seat to our array
		seats = append(seats, seat_id)
	}
	fmt.Println("max seat", max_seat)

	// Sort our list of seat ids
	sort.Slice(seats, func(i, j int) bool {
		return seats[i] < seats[j]
	})

	// Find the free seat

	for i := 1; i < len(seats); i++ {
		if seats[i-1]+2 == seats[i] {
			fmt.Println("My seat is", seats[i]-1)
			break
		}
	}
	//	fmt.Println(seats)
}
