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
	"os"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Moodly
//
//   1. Get username from command-line
//
//   2. Display the usage if the username is missing
//
//   3. Create an array
//     1. Add three positive mood messages
//     2. Add three negative mood messages
//
//   4. Randomly select and print one of the mood messages
//
// EXPECTED OUTPUT
//
//   go run main.go
//     [your name]
//
//   go run main.go Socrates
//     Socrates feels good 👍
//
//   go run main.go Socrates
//     Socrates feels bad 👎
//
//   go run main.go Socrates
//     Socrates feels sad 😞
//
//   go run main.go Socrates
//     Socrates feels happy 😀
//
//   go run main.go Socrates
//     Socrates feels awesome 😎
//
//   go run main.go Socrates
//     Socrates feels terrible 😩
// ---------------------------------------------------------

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Print("go run main.go [your name]")
		return
	}

	messages := []string{
		"feels good 👍",
		"feels bad 👎",
		"feels sad 😞",
		"feels happy 😀",
		"feels awesome 😎",
		"feels terrible 😩",
	}

	fmt.Print(len(messages))
	rand.Seed(time.Now().UnixNano())
	indexMessage := rand.Intn(len(messages))

	fmt.Printf("%s %s", args[0], messages[indexMessage])
}
