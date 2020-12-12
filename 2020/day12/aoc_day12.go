// https://adventofcode.com/2020/day12
//
// Read the file from stdin

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var lat = 0  // (+)north/south
var long = 0 // (+)east/west
var dir = 90

func pt1_rule(opcode string, value int) {
	switch opcode {
	case "N":
		lat += value
	case "S":
		lat -= value
	case "E":
		long += value
	case "W":
		long -= value
	case "F":
		switch dir {
		case 90:
			long += value
		case 180:
			lat -= value
		case 270:
			long -= value
		case 0:
			lat += value
		default:
			fmt.Println("Unknown degrees", value)
		}
	case "L":
		dir -= value
		dir %= 360
		if dir < 0 {
			dir += 360
		}
	case "R":
		dir += value
		dir %= 360
	default:
		fmt.Println("Unknown command", opcode)
	}
	// fmt.Println(opcode, value, lat, long, dir)
}

var waypt_lat = 1
var waypt_long = 10
var ship_lat = 0
var ship_long = 0

func pt2_rule(opcode string, value int) {
	switch opcode {
	case "N":
		waypt_lat += value
	case "S":
		waypt_lat -= value
	case "E":
		waypt_long += value
	case "W":
		waypt_long -= value
	case "F":
		ship_lat += value * waypt_lat
		ship_long += value * waypt_long
	case "L":
		n_long := 0
		n_lat := 0
		switch value {
		case 90:
			n_long = -waypt_lat
			n_lat = waypt_long
		case 180:
			n_long = -waypt_long
			n_lat = -waypt_lat
		case 270:
			n_long = waypt_lat
			n_lat = -waypt_long
		}
		waypt_lat = n_lat
		waypt_long = n_long
	case "R":
		n_long := 0
		n_lat := 0
		switch value {
		case 90:
			n_long = waypt_lat
			n_lat = -waypt_long
		case 180:
			n_long = -waypt_long
			n_lat = -waypt_lat
		case 270:
			n_long = -waypt_lat
			n_lat = waypt_long
		}
		waypt_lat = n_lat
		waypt_long = n_long
	default:
		fmt.Println("Unknown command", opcode)
	}
	// fmt.Println(opcode, value, waypt_lat, waypt_long, ship_lat, ship_long)
}

func main() {

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		opcode := line[0:1]
		value, _ := strconv.Atoi(line[1:]) // all values are positive

		pt1_rule(opcode, value)
		pt2_rule(opcode, value)
	}

	fmt.Println("Manhattan distance #1 (", lat, long, ")", Abs(lat)+Abs(long))
	fmt.Println("Manhattan distance #2 (", ship_lat, ship_long, ")", Abs(ship_lat)+Abs(ship_long))
}
