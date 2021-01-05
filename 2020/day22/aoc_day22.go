// https://adventofcode.com/2020/day22
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

type deck = [2][]int

var pt2 bool

func init() {
	const (
		defaultV = false
		usage    = "enable pt2"
	)

	flag.BoolVar(&pt2, "2", defaultV, usage)

}

// returns winner (player1 == 0 ; player2 == 1)
func play_game2(cards deck, game int) int {
	game_history := []deck{}
	// fmt.Println("game", game, "cards", cards)

	round := 1
	for true {
		var winner int
		var loser int

		var top_card [2]int

		// Check for infinite recursion
		for k, c := range game_history {
			// Don't test the most recent entry as it WILL match - we just added it on the previous round
			if k == len(game_history)-1 {
				break
			}

			if len(c[0]) != len(cards[0]) || len(c[1]) != len(cards[1]) {
				continue
			}
			m_req := len(c[0]) + len(c[1])
			m := 0
			for kk, cc := range c {
				for kkk, ccc := range cc {
					if ccc == cards[kk][kkk] {
						m++
					}
				}
			}
			if m_req == m {
				// fmt.Println("recursion detection - player 1 wins", "history", game_history, "cards", cards)
				return 0
			}
		}

		top_card[0] = cards[0][0]
		top_card[1] = cards[1][0]
		cards[0] = cards[0][1:]
		cards[1] = cards[1][1:]

		// fmt.Println("game", game, "round", round, "player1 plays", top_card[0], "player2 plays", top_card[1])

		if (len(cards[0]) >= top_card[0]) && (len(cards[1]) >= top_card[1]) {
			// recurse to find the winner
			var new_cards deck
			for j, _ := range top_card {
				for i := 0; i < top_card[j]; i++ {
					new_cards[j] = append(new_cards[j], cards[j][i])
				}
			}
			winner = play_game2(new_cards, game+1)
			if winner == 1 {
				loser = 0
			} else {
				loser = 1
			}

		} else {
			if top_card[0] > top_card[1] {
				winner = 0
				loser = 1
			} else {
				winner = 1
				loser = 0
			}
		}

		// fmt.Println("game", game, "round", round, "winner player", winner+1)

		// Move cards
		cards[winner] = append(cards[winner], top_card[winner])
		cards[winner] = append(cards[winner], top_card[loser])

		// fmt.Println("cards", cards)

		game_history = append(game_history, cards)

		if len(cards[loser]) == 0 {

			if game == 1 {
				// Calc score
				weight := len(cards[winner])
				score := 0
				for _, v := range cards[winner] {
					score += v * weight
					weight -= 1

				}
				fmt.Println("Score", score)
			}
			return winner
			break
		}
		round++
	}
	return 0
}

func play_game1(cards deck) {
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

func main() {

	flag.Parse()

	var cards = deck{}

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
	if pt2 {
		winner := play_game2(cards, 1)
		fmt.Println("winner", winner)

	} else {
		play_game1(cards)
	}
}
