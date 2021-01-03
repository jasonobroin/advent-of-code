// https://adventofcode.com/2020/day22
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

func main() {

	var cards = [2][]int{}

	// Parse the input stream to parse each line
	scanner := bufio.NewScanner(os.Stdin)

	// Ticket definitions

	player := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Player") {
			continue
		}

		if len(line) == 0 {
			player++
			continue
		}

		card, _ := strconv.Atoi(line)
		cards[player] = append(cards[player], card)
	}

	fmt.Println("cards", cards)

	// Play game

	round := 1
	for true {
		var winner int
		var loser int

		if cards[0][0] > cards[1][0] {
			winner = 0
			loser = 1
		} else {
			winner = 1
			loser = 0
		}

		// fmt.Println("round", round, "player1 plays", cards[0][0], "player2 plays", cards[1][0], "winner player", winner+1)

		// Move cards
		cards[winner] = append(cards[winner], cards[winner][0])
		cards[winner] = append(cards[winner], cards[loser][0])
		cards[winner] = cards[winner][1:]
		cards[loser] = cards[loser][1:]

		// fmt.Println("cards", cards)

		if len(cards[loser]) == 0 {
			// fmt.Println("Player", loser+1, "lost")

			// Calc score
			weight := len(cards[winner])
			score := 0
			for _, v := range cards[winner] {
				score += v * weight
				weight -= 1

			}
			fmt.Println("Score", score)
			break
		}
		round++
	}
}
