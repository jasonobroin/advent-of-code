// https://adventofcode.com/2020/day19
//
// Read the file from stdin

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var pt2 bool

func init() {
	const (
		defaultV = false
		usage    = "enable pt2"
	)

	flag.BoolVar(&pt2, "2", defaultV, usage)

}

type rule struct {
	match    string
	subrule1 []string
	subrule2 []string
}

var rules = map[string]rule{}

// Returning <true> when there is no match allows more cases to be valid, but
// this case is incorrectly marked valid
// aaaabbaaaabbaaa
//
// However, return <true> when there is no character to check seems wrong - that seems
// like an failure and we should return false in that case; however, we miss a number
// of valid cases
//
// I think this is the question - why do we fail valid cases when we return
// false from check_valid with no string?
//
// Presumably there is some path where we should return true in this case, or not
// prematurely return false - maybe there's a path we've not explored completely
// so we return up the stack too quickly?

func check_valid(s string, rule string) (bool, int) {
	if s == "" {
		// Consumed all characters - nothing left to check
		// fmt.Println(rule, "empty")
		return true, 0
	}

	if rules[rule].match != "" {
		// We have a character to check
		if rules[rule].match == string(s[0]) {
			// fmt.Println(rule, "match", rules[rule].match)
			return true, 1
		} else {
			return false, 0
		}
	} else {
		// Check both branches
		rule1_good := false
		rule2_good := false
		off := 0
		o1 := 0
		o2 := 0
		for _, v := range rules[rule].subrule1 {
			// fmt.Println(rule, "check1", s[o1:], "against", v)
			rule1_good, off = check_valid(s[o1:], v)
			if !rule1_good {
				break
			}
			o1 += off
		}
		for _, v := range rules[rule].subrule2 {
			// fmt.Println(rule, "check2", s[o2:], "against", v)
			rule2_good, off = check_valid(s[o2:], v)
			if !rule2_good {
				break
			}
			o2 += off
		}

		if rule1_good && rule2_good {
			fmt.Println(rule1_good, rule2_good)
		}

		if rule1_good {
			return true, o1
		}
		if rule2_good {
			return true, o2
		}
		return false, o2
	}

	return false, 0
}

func main() {

	flag.Parse()

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)

	// Ticket definitions

	for scanner.Scan() {
		line := scanner.Text()

		if pt2 {
			if line == "8: 42" {
				line += " | 42 8"
			}
			if line == "11: 42 31" {
				line += " | 42 11 31"
			}
		}

		if len(line) == 0 {
			break
		}

		first := strings.Split(line, ":")
		rule_num := first[0]
		if strings.Contains(first[1], "\"") {
			sub := strings.Split(first[1], "\"")
			fmt.Println(rule_num, sub[1])
			rules[rule_num] = rule{sub[1], nil, nil}

		} else {

			second := strings.Split(first[1], "|")
			subrule1 := strings.Fields(second[0])
			var subrule2 []string
			if len(second) == 2 {
				subrule2 = strings.Fields(second[1])
			}
			fmt.Println(rule_num, subrule1, subrule2)
			rules[rule_num] = rule{"", subrule1, subrule2}
		}
	}

	fmt.Println(rules)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		if ok, off := check_valid(line, "0"); ok {
			if off == len(line) {
				fmt.Println("valid", line, off)
				total++
			} else {
				fmt.Println("valid up to a point", line, off)
			}
		} else {
			fmt.Println("invalid", line, off)
		}
	}

	fmt.Println("Number of matching messages", total)
}
