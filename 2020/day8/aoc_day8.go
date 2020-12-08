// https://adventofcode.com/2020/day8
//
// Read the file from stdin

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type inst struct {
	opcode  string
	val     int
	visited bool
}

func run_program(program []inst) bool {
	// Run program

	pc := 0
	acc := 0

	prev_jmp := 0
	for true {
		// Check PC in range first
		if pc >= len(program) {
			fmt.Println("correct ending. acc=", acc)
			return true
		}
		if program[pc].visited {
			fmt.Println("infinite loop. acc=", acc)
			return false
		}
		program[pc].visited = true
		switch program[pc].opcode {
		case "acc":
			acc += program[pc].val
			pc++
		case "nop":
			pc++
		case "jmp":
			if pc > prev_jmp {
				prev_jmp = pc
			}
			pc += program[pc].val
		}
	}

	// never gets here
	return false

}

func main() {

	var program []inst

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		opcode := line[0:3]
		value, _ := strconv.Atoi(line[4:])
		//		fmt.Println(opcode, value)
		program = append(program, inst{opcode, value, false})
	}

	fmt.Println("Part 1")
	run_program(program)

	fmt.Println("Part 2")
	// Find the bad instruction
	for i, _ := range program {
		// reset program state
		for k, _ := range program {
			program[k].visited = false
		}
		old := program[i].opcode
		program[i].opcode = "nop"
		if !run_program(program) {
			program[i].opcode = old
		} else {
			fmt.Println("Fixed instruction", i)
			break
		}
	}
}
