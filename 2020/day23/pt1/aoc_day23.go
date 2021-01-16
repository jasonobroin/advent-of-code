// https://adventofcode.com/2020/day23
//
// Read the file from stdin

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func remove(line string, current_cup int, num int) (string, string) {
	removed := ""
	new_str := ""

	offset := (current_cup + 1) % len(line)
	atend := len(line) - offset

	if atend < num {
		removed = line[offset:]
		removed += line[:num-atend]
		new_str = line[num-atend : offset]
	} else {
		removed = line[offset : offset+num]
		new_str = line[:offset]
		new_str += line[offset+num:]

	}

	return new_str, removed
}

func new_label(line string, label int) int {
	min := 9
	max := 0
	for _, v := range line {
		vv, _ := strconv.Atoi(string(v))
		if vv > max {
			max = vv
		}
		if vv < min {
			min = vv
		}
	}
	for true {
		label--
		if label < min {
			return max
		}
		test := strconv.Itoa(label)
		if strings.Contains(line, test) {
			return label
		}
	}
	return 0
}

func add(line string, dest_cup int, new_cups string) string {
	if dest_cup > len(line) {
		dest_cup = len(line)
	}
	new_str := line[:dest_cup] + new_cups
	if len(line) > dest_cup {
		new_str += line[dest_cup:]
	}
	return new_str
}

func find_loc(line string, label int) int {
	for k, v := range line {
		vv, _ := strconv.Atoi(string(v))
		if label == vv {
			return k
		}
	}
	return 0
}

func final_line(line string) string {
	for k, v := range line {
		vv, _ := strconv.Atoi(string(v))
		if 1 == vv {
			new_str := line[k+1:]
			new_str += line[:k]
			return new_str
		}
	}
	return ""
}

func main() {

	flag.Parse()

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		current_cup := 0
		for i := 0; i < count; i++ {
			fmt.Println("round", i+1, current_cup, "\"", string(line[current_cup]), "\"", line)
			label, _ := strconv.Atoi(string(line[current_cup]))
			current_label := label
			pickup := ""
			line, pickup = remove(line, current_cup, 3)
			// fmt.Println("pickup", pickup, line, label)
			label = new_label(line, label)
			loc := find_loc(line, label)
			// fmt.Println("dest cup", label, "loc", loc)
			line = add(line, loc+1, pickup)
			current_cup = find_loc(line, current_label)
			current_cup += 1
			current_cup %= len(line)
		}
		fmt.Println("final line", line, final_line(line))

	}
}
