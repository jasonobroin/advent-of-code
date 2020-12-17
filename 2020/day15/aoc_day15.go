// https://adventofcode.com/2020/day13
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

var target int

func init() {
	const (
		defaultTarget = 2020
		usage         = "target entry size"
	)

	flag.IntVar(&target, "t", defaultTarget, usage)

}

type info struct {
	last_turn int
	prev_turn int
}

var spoken = map[int]info{}
var turn int = 1

func update_info(num int, turn int) bool {
	i, present := spoken[num]
	if present {
		i.prev_turn = i.last_turn
	} else {
		i.prev_turn = 0
	}
	i.last_turn = turn
	spoken[num] = i

	return present
}

func main() {

	flag.Parse()

	fmt.Println("Target =", target)

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	input := strings.Split(line, ",")
	fmt.Println(input)

	last := 0
	var present bool = false

	for _, v := range input {
		val, _ := strconv.Atoi(v)
		present = update_info(val, turn)
		last = val
		turn++
	}

	val := 0
	for turn <= target {
		if !present {
			val = 0
			present = update_info(val, turn)
			// Special case - 0 may not have been added in the initial data set;
			// we need to ensure both its last_turn and prev_turn are set to the
			// current turn value
			if !present {
				_ = update_info(val, turn)
			}
			last = 0
			present = true
		} else {
			i, _ := spoken[last]
			val := i.last_turn - i.prev_turn
			present = update_info(val, turn)
			last = val
		}
		turn++
	}
	fmt.Println("Last number", last)
}
