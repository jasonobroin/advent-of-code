// https://adventofcode.com/2020/day7
//
// Iterate to find total possible bags that can contain a specific bag (shiny gold)
//
// rule grammar:
// <bag type> bags contain [<n> <bag type> bag(s),]
//
// <n> can be a number or "no other"
//
// bags that contain "no other" can be completely ignored
//
// Search - for each bag type, recursively search each bag type inside it to see if
// it contains the search bag type. Count the total we find. Once we find our search
// bag we can exit the current search
//
// The search could be optimized if we record any time we know that a bag contains our
// search bag (i.e. we have already descended this path)
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

type bag struct {
	num  int
	name string
}

type bag_contents struct {
	bags []bag
}

var all_bags = map[string]bag_contents{}

func subcount_bags(s string, f string) int {
	bags := all_bags[s]
	for _, v1 := range bags.bags {
		if v1.name == f {
			//			fmt.Println("found bag in", s)
			return 1
		} else {
			if subcount_bags(v1.name, f) == 1 {
				return 1
			}
		}
	}
	return 0
}

func count_bags(s string) int {
	found := 0
	for k, _ := range all_bags {
		found += subcount_bags(k, s)
	}
	return found
}

func count_bags_inside(s string) int {
	search_bag := all_bags[s]

	count := 0

	for _, v := range search_bag.bags {
		count += v.num * (1 + count_bags_inside(v.name))
	}
	return count
}

func main() {

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Decode bag info

		if strings.Contains(line, "no other") {
			fmt.Println("Skip", line)
			continue
		}

		first := strings.Split(line, "bags contain ")
		bag_type := strings.TrimSpace(first[0])

		// Split the bag contents
		bags_within := strings.Split(first[1], ",")

		var bag_cont bag_contents

		for _, b := range bags_within {
			st := strings.TrimSpace(b)
			s := strings.SplitN(st, " ", 2)
			num, _ := strconv.Atoi(s[0])
			bn := strings.Split(s[1], "bag")
			bag_name := strings.TrimSpace(bn[0])

			bag_cont.bags = append(bag_cont.bags, bag{num, bag_name})
		}
		all_bags[bag_type] = bag_cont
	}

	// Part1 - how many bags can hold a shiny gold bag

	total := count_bags("shiny gold")
	fmt.Println("total found", total)

	// Part2 - How many bags do we have inside our shiny gold bag

	total2 := count_bags_inside("shiny gold")
	fmt.Println("total pt2 ", total2)
}
