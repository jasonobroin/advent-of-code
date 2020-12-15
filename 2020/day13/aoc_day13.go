// https://adventofcode.com/2020/day13
//
// Read the file from stdin

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type bus_info struct {
	offset  int64
	bus_num int64
}

var one = big.NewInt(1)

func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

// Using Chinese Remainer Theorem from https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
func find_bus_seq(lines []string) {

	var n = []*big.Int{}
	var a = []*big.Int{}

	for k, v := range lines {
		if v == "x" {
			continue
		}
		bus_num, _ := strconv.ParseInt(v, 10, 64)
		n = append(n, big.NewInt(bus_num))
		a = append(a, big.NewInt(-int64(k)))
	}

	res, _ := crt(a, n)
	fmt.Println("first bus @", res)
}

func find_first_bus(first_time int64, lines []string) {

	var bus_line int64 = 0
	var earliest_time int64 = 100000000 // a large number
	for _, v := range lines {
		if v == "x" {
			continue
		}
		bus_num, _ := strconv.ParseInt(v, 10, 64)
		last_bus_time := int64(first_time/bus_num) * bus_num
		//		fmt.Println(v, last_bus_time)
		if last_bus_time+bus_num < earliest_time {
			earliest_time = last_bus_time + bus_num
			bus_line = bus_num
		}
	}
	fmt.Println("bus line", bus_line, "at", earliest_time)
	fmt.Println("result = ", (earliest_time-first_time)*bus_line)
}

func main() {

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line1 := scanner.Text()
	scanner.Scan()
	line2 := scanner.Text()
	first_time, _ := strconv.ParseInt(line1, 10, 64)
	lines := strings.Split(line2, ",")

	find_first_bus(first_time, lines)
	find_bus_seq(lines)
}
