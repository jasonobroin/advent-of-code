// https://adventofcode.com/2020/day18
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

func eval(l []string) int {
	var stack []int

	for _, v := range l {
		switch v {
		case "+":
			{
				n := len(stack) - 1
				f := stack[n]
				s := stack[n-1]
				sum := f + s
				stack = stack[:n-1]
				stack = append(stack, sum)
			}
		case "*":
			{
				n := len(stack) - 1
				f := stack[n]
				s := stack[n-1]
				sum := f * s
				stack = stack[:n-1]
				stack = append(stack, sum)

			}
		default:
			f, _ := strconv.Atoi(v)
			stack = append(stack, f)

		}
	}
	sum := stack[0]
	return sum
}

// Part1 translate into RPN. L->R ordering, with ()

func translate1(l string) ([]string, int) {

	var rpn []string

	start_offset := 0
	end_offset := len(l)
	first := true
	op := ""
	fop := ""
	end := false

	for true {
		offset := strings.IndexAny(l[start_offset:end_offset], "+*()")
		if offset == -1 {
			op = l[start_offset:]
			end = true
		}
		// Do we have an operator? No... a number
		if offset > 0 {
			op = l[start_offset : start_offset+offset]
			start_offset += offset
		} else {
			op = l[start_offset : start_offset+1]
			start_offset += 1
		}
		switch op {
		case ")":
			return rpn, start_offset
		case "+":
			fop = op
		case "*":
			fop = op
		case "(":
			r, offset := translate1(l[start_offset:])
			for _, v := range r {
				rpn = append(rpn, v)
			}
			if fop != "" {
				rpn = append(rpn, fop)
			}
			start_offset += offset
			if start_offset >= end_offset {
				return rpn, start_offset
			}
			first = false
		default:
			rpn = append(rpn, op)
			if !first {
				rpn = append(rpn, fop)
			} else {
			}
			fop = op
			first = false

		}
		if end {
			break
		}
	}
	return rpn, start_offset
}

// Part2 translate into RPN. L->R ordering, with (); + has higher priority than *

func translate2(l string, rtn_op string) ([]string, int) {
	var rpn []string

	start_offset := 0
	end_offset := len(l)
	first := true
	op := ""
	fop := ""
	end := false

	for true {
		offset := strings.IndexAny(l[start_offset:end_offset], "+*()")
		if offset == -1 {
			op = l[start_offset:]
			end = true
		}
		// Do we have an operator? No... a number
		if offset > 0 {
			op = l[start_offset : start_offset+offset]
			start_offset += offset
		} else {
			op = l[start_offset : start_offset+1]
			start_offset += 1
		}
		switch op {
		case ")":
			if rtn_op == "*" {
				start_offset -= 1
			}
			return rpn, start_offset
		case "+":
			fop = op
		case "*":
			// We're inside  scope - seeing another * indicates the end of the scope, so return with * unprocesed
			if rtn_op == "*" {
				return rpn, start_offset - 1
			}
			fop = op
			// * has lower precedence than other operators, so treat like parentheses with * at the end of the scope
			r, offset := translate2(l[start_offset:], "*")
			for _, v := range r {
				rpn = append(rpn, v)
			}
			if fop != "" {
				rpn = append(rpn, fop)
			}
			start_offset += offset
			if start_offset >= end_offset {
				return rpn, start_offset
			}
			first = false
			// We need to see if the next operator is a +
		case "(":
			r, offset := translate2(l[start_offset:], "")
			// offset -= 1
			for _, v := range r {
				rpn = append(rpn, v)
			}
			if fop != "" {
				rpn = append(rpn, fop)
			}
			start_offset += offset
			if start_offset >= end_offset {
				return rpn, start_offset
			}
			first = false
		default:
			rpn = append(rpn, op)
			if !first {
				rpn = append(rpn, fop)
			} else {
			}
			fop = op
			first = false

		}
		if end {
			break
		}
	}
	return rpn, start_offset
}

func main() {

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)

	// Ticket definitions

	total1 := 0
	total2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		a, _ := translate1(strings.ReplaceAll(line, " ", ""))
		res := eval(a)
		total1 += res

		b, _ := translate2(strings.ReplaceAll(line, " ", ""), "")
		res2 := eval(b)
		total2 += res2

	}
	fmt.Println("total1", total1)
	fmt.Println("total2", total2)
}
