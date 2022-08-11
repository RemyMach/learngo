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
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

func main() {
	var (
		min      int
		errorMin error
		max      int
		errorMax error
	)

	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Printf("you have to fill min and max args")
		return
	}

	if min, errorMin = strconv.Atoi(args[0]); errorMin != nil {
		fmt.Printf("the first arg is not a number")
		return
	}

	if max, errorMax = strconv.Atoi(args[1]); errorMax != nil {
		fmt.Printf("the secund arg is not a number")
		return
	}

	if min >= max {
		fmt.Printf("first arg have to be lower than max")
		return
	}

	var sum int

	for i := min; i <= max; i++ {
		sum += i
		fmt.Printf("%d", i)
		if i < max {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d", sum)
}
