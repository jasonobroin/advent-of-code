// https://adventofcode.com/2020/day16
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

type void struct{}

var member void
var labels = map[string]void{}

type scope struct {
	start int
	end   int
}

type fields = []scope

var tickets = map[string]fields{}

type set_labels struct {
	done  bool
	label map[string]void
}

func add_range(t string, r []string) {

	var scopes fields

	for _, v := range r {
		val := strings.Split(v, "-")
		start, _ := strconv.Atoi(strings.TrimSpace(val[0]))
		end, _ := strconv.Atoi(strings.TrimSpace(val[1]))
		s := scope{start, end}
		scopes = append(scopes, s)
	}
	tickets[t] = scopes

	labels[t] = member

}

// Check if val is in one our ticket ranges

func check_val(val int) bool {
	for _, v := range tickets {
		for _, vv := range v {
			if val >= vv.start && val <= vv.end {
				return true
			}
		}
	}
	return false
}

func main() {

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)

	// Ticket definitions

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		first := strings.Split(line, ":")
		ranges := strings.Split(first[1], " or ")

		add_range(first[0], ranges)

	}

	// "Your ticket"

	scanner.Scan()
	line := scanner.Text() // "your ticket:"
	scanner.Scan()
	line = scanner.Text()
	my_tickets := strings.Split(line, ",")
	scanner.Scan()
	line = scanner.Text() // empty line
	scanner.Scan()
	line = scanner.Text() // "nearby ticket:"

	fmt.Println("my_tickets", my_tickets)

	// nearby tickets

	error_rate := 0

	// General approach - for each column in the nearby tickets, start with an array of all possible labels
	// As we look up a value, return all labels it is NOT found it - this means check_val needs to change to
	// continue searching
	// Remove all labels NOT found from the array of possible labels for this column
	//
	// We could abstract our labels - ultimately we need something that represents the label, but we're not optimizing for speed, so the string will do fine
	//
	// At the end, there should only be one label left for each row

	num_labels := len(my_tickets)
	var final_labels []set_labels
	for k := 0; k < num_labels; k++ {
		new_set := set_labels{}
		new_set.label = make(map[string]void)
		for a, b := range labels {
			new_set.label[a] = b
		}
		final_labels = append(final_labels, new_set)
	}

	for scanner.Scan() {
		line := scanner.Text()

		ranges := strings.Split(line, ",")

		good_ticket := true
		for _, v := range ranges {
			val, _ := strconv.Atoi(v)
			ok := check_val(val)
			if !ok {
				error_rate += val
				good_ticket = false
			}
		}

		if good_ticket {
			for k, v := range ranges {
				val, _ := strconv.Atoi(v)

				for kk, vv := range tickets {
					found := false
					for _, vvv := range vv {
						if val >= vvv.start && val <= vvv.end {
							found = true
						}
					}
					if !found {
						// Remove this ticket name
						delete(final_labels[k].label, kk)
					}
				}
			}
		}
	}

	// Finally as we determine which column is what, remove that label in the rest of the columns
	// We need to do this iteratively; removing the name of the column with only one entry on each pass
	// until we have reduced all columns to a single entry

	for true {
		removed := false
		for k, v := range final_labels {
			if v.done {
				continue
			}

			if len(v.label) == 1 && !v.done {
				// For each entry in the current label, remove from other labels
				for i, _ := range v.label {
					for kk := 0; kk < len(final_labels); kk++ {
						if final_labels[kk].done {
							continue
						}
						if kk == k {
							continue
						}
						delete(final_labels[kk].label, i)
						removed = true
					}
				}
				final_labels[k].done = true
			}
		}
		if !removed {
			break
		}

	}

	mult := 1
	for k, v := range final_labels {
		for i, _ := range v.label {
			fmt.Println("My ticket", i, "=", my_tickets[k])
			if strings.Contains(i, "departure") {
				tick_val, _ := strconv.Atoi(my_tickets[k])
				mult *= tick_val
			}
		}
	}

	fmt.Println(error_rate)
	fmt.Println("Departure mult =", mult)
}
