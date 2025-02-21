// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

func main() {

	guess := 6
	turn := 0
	for {
		rand.Seed(time.Now().UnixNano())

		value := rand.Intn(10)

		if (turn == 0 && value == guess) {
			fmt.Printf("Congratulations you are the best")
			return
		} else if value == guess {
			fmt.Printf("Congratulations you find at turn %d", turn+1)
			return
		}

		turn++
	}
}
