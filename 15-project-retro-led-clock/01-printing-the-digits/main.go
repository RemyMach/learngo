// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {

	zero := []string{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := []string{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := []string{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := []string{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := []string{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := []string{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := []string{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := []string{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := []string{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := []string{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	digits := [10][]string{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	var printableDigits [5]string

	for _, num := range digits {
		for index, line := range num {
			fmt.Printf("%d\n", index)
			printableDigits[index] += " " + line
		}
	}

	for _, line := range printableDigits {
		fmt.Printf("%s\n", line)
	}

}
