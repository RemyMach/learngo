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
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

func main() {

	sentencesWin := []string{"YOU WON", "YOU'RE AWESOME"}
	sentencesLost := []string{"LOSER!", "YOU LOST. TRY AGAIN?"}
	rand.Seed(time.Now().UnixNano())
	indexRandom := rand.Intn(2)
	winOrLost := rand.Intn(2)
	switch winOrLost {
	case 0:
		fmt.Printf("%s", sentencesWin[indexRandom])
	case 1:
		fmt.Printf("%s", sentencesLost[indexRandom])

	}

}
