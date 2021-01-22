// https://adventofcode.com/2020/day23
//
// Read the file from stdin
//

package main

import (
	"bufio"
	"container/ring"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
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

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func new_label(r *ring.Ring, s *ring.Ring, label int, max int) int {
	min := 0

	for true {
		found := false
		s.Do(func(p interface{}) {
			if p.(int) == min {
				min++
				found = true
			}
		})
		if !found {
			break
		}
	}

	for true {
		found := false
		s.Do(func(p interface{}) {
			if p.(int) == max {
				max--
				found = true
			}
		})
		if !found {
			break
		}
	}

	for true {
		label--
		found := false
		s.Do(func(p interface{}) {
			if p.(int) == label {
				found = true
			}
		})
		if !found {
			break
		}
	}

	if label <= min {
		return max
	}
	return label
}

func two_cups(r *ring.Ring) {
	for true {
		if 1 == r.Value.(int) {
			r = r.Next()
			// fmt.Println(r.Value.(int))
			res := int64(r.Value.(int))
			r = r.Next()
			// fmt.Println(r.Value.(int))
			res *= int64(r.Value.(int))
			fmt.Println("Product = ", res)
			break
		} else {
			r = r.Next()
		}
	}
}

func printRing(r *ring.Ring) {
	r.Do(func(p interface{}) {
		fmt.Print(p.(int), " ")
	})
	fmt.Print("\n")
}

func printFinal(r *ring.Ring) {
	for true {
		if r.Value.(int) == 1 {
			s := ""
			r.Do(func(p interface{}) {
				s += strconv.Itoa(p.(int))
			})
			fmt.Println("final =", s[1:])
			break
		} else {
			r = r.Next()
		}
	}
}

func main() {

	flag.Parse()

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		allocate := len(line)
		if pt2 {
			allocate = 1000000
		}

		// Build a map of values -> ring entries; this should allow us to
		// very quickly find an entry we care about (at a cost in memory)
		var lookup = map[int]*ring.Ring{}

		r := ring.New(allocate)
		for _, v := range line {
			val, _ := strconv.Atoi(string(v))
			r.Value = val
			lookup[val] = r
			r = r.Next()
		}

		if pt2 {
			for i := len(line) + 1; i <= allocate; i++ {
				r.Value = i
				lookup[i] = r
				r = r.Next()
			}
		}

		for i := 0; i < count; i++ {
			label := r.Value
			pickup := r.Unlink(3)

			label = new_label(r, pickup, label.(int), allocate)

			l := lookup[label.(int)]
			_ = l.Link(pickup)

			r = r.Next()

		}
		if pt2 {
			two_cups(r)
		} else {
			printFinal(r)
		}
	}
}
