package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type player struct {
	number    int
	sign      string
	opponents []int
}

var signs = map[string][]string{
	"R": {"L", "C"},
	"P": {"R", "S"},
	"C": {"L", "P"},
	"L": {"S", "P"},
	"S": {"C", "R"},
}

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var numberOfParticipants int
	fmt.Scan(&numberOfParticipants)

	var players []player

	// initilise players
	for i := 0; i < numberOfParticipants; i++ {
		var playerNum int
		var playerSign string
		fmt.Scan(&playerNum, &playerSign)

		players = append(players, player{number: playerNum, sign: playerSign})
	}

	// play games
	// calculate num of rounds
	roundsLeft := int(math.Floor(math.Log2(float64(numberOfParticipants))))

	for round := roundsLeft; round > 0; round-- {
		matchesThisRound := len(players) / 2
		// for each game in round
		currentMatch := 0

		var roundLosers []int

		for match := matchesThisRound; match > 0; match -= 1 {
			winner := getWinner(players[currentMatch], players[currentMatch+1])

			// Update each player with who they played
			players[currentMatch].opponents = append(players[currentMatch].opponents, players[currentMatch+1].number)
			players[currentMatch+1].opponents = append(players[currentMatch+1].opponents, players[currentMatch].number)

			if winner == 0 {
				roundLosers = append(roundLosers, currentMatch+1)
			} else {
				roundLosers = append(roundLosers, currentMatch)
			}

			currentMatch += 2
			matchesThisRound--
		}

		// remove losers
		sort.Slice(roundLosers, func(i, j int) bool {
			return roundLosers[i] > roundLosers[j]
		})

		for _, index := range roundLosers {
			players = append(players[:index], players[index+1:]...)
		}

		roundsLeft--
	}

	fmt.Println(players[0].number)
	opponents := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(players[0].opponents)), " "), "[]")
	fmt.Println(opponents)
}

func getWinner(player1, player2 player) int {
	// If they draw
	if player1.sign == player2.sign {
		if player1.number < player2.number {
			return 0
		} else {
			return 1
		}
	}

	player1CanBeat := signs[player1.sign]

	if contains(player1CanBeat, player2.sign) {
		return 0
	} else {
		return 1
	}
}

func contains(signs []string, sign string) bool {
	for _, a := range signs {
		if a == sign {
			return true
		}
	}
	return false
}
