// https://adventofcode.com/2020/day23
//
// Read the file from stdin
//
// part 2
//
// horribly slow - try using container/ring

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var pt2 bool
var count int

func init() {
	const (
		defaultV = false
		usage    = "enable pt2"
	)

	flag.BoolVar(&pt2, "2", defaultV, usage)

}

func init() {
	const (
		defaultV = 100
		usage    = "count of rounds"
	)

	flag.IntVar(&count, "c", defaultV, usage)

}

type array = []int

func remove(line array, current_cup int, num int) (array, array) {
	var removed array
	var new_str array

	offset := (current_cup + 1) % len(line)
	atend := len(line) - offset

	if atend < num {
		removed = append(removed, line[offset:]...)
		removed = append(removed, line[:num-atend]...)
		new_str = append(new_str, line[num-atend:offset]...)
	} else {
		removed = append(removed, line[offset:offset+num]...)
		new_str = append(new_str, line[:offset]...)
		new_str = append(new_str, line[offset+num:]...)

	}

	return new_str, removed
}

func new_label(line array, label int) int {
	min := 9
	max := 0
	for _, v := range line {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	for true {
		label--
		if label < min {
			return max
		}
		for _, v := range line {
			if v == label {
				return label
			}
		}
	}
	return 0
}

func add(line array, dest_cup int, new_cups array) array {
	if dest_cup > len(line) {
		dest_cup = len(line)
	}
	var new_str array
	new_str = append(new_str, line[:dest_cup]...)
	new_str = append(new_str, new_cups...)
	if len(line) > dest_cup {
		new_str = append(new_str, line[dest_cup:]...)
	}
	return new_str
}

func find_loc(line array, label int) int {
	for k, v := range line {
		if label == v {
			return k
		}
	}
	return 0
}

func final_line(line array) string {
	for k, v := range line {
		if 1 == v {
			new_str := line[k+1:]
			new_str = append(new_str, line[:k]...)

			s := ""
			for _, v := range new_str {
				s += strconv.Itoa(v)
			}
			return s
		}
	}
	return ""
}

func two_cups(line array) {
	for k, v := range line {
		if 1 == v {
			fmt.Println("Found 1 @ ", k, len(line))
			next := k + 1
			if next >= len(line) {
				next = 0
			}
			fmt.Println(line[next])
			res := int64(line[next])
			next++
			if next >= len(line) {
				next = 0
			}
			fmt.Println(line[next])
			res *= int64(line[next])
			fmt.Println("Product = ", res)
		}
	}
}

func main() {

	flag.Parse()

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		var carray array
		for _, v := range line {
			val, _ := strconv.Atoi(string(v))
			carray = append(carray, val)
		}
		fmt.Println(carray)

		if pt2 {
			for i := 10; i <= 1000000; i++ {
				carray = append(carray, i)
			}
		}

		current_cup := 0
		for i := 0; i < count; i++ {
			// fmt.Println("round", i+1, current_cup, "\"", carray[current_cup], "\"")
			label := carray[current_cup]
			current_label := label
			var pickup array
			carray, pickup = remove(carray, current_cup, 3)
			// fmt.Println("pickup", pickup, carray, label)
			label = new_label(carray, label)
			loc := find_loc(carray, label)
			// fmt.Println("dest cup", label, "loc", loc)
			carray = add(carray, loc+1, pickup)
			current_cup = find_loc(carray, current_label)
			current_cup += 1
			current_cup %= len(carray)
			if i%100000 == 0 {
				fmt.Println("round", i)
			}
		}
		if pt2 {
			two_cups(carray)
		} else {
			fmt.Println("final line", carray, final_line(carray))
		}
	}
}
