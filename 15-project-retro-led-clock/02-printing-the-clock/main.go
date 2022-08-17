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
	"time"
)

func main() {
	type placeholder [5]string

	zero := placeholder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}

	one := placeholder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}

	two := placeholder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}

	three := placeholder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}

	four := placeholder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}

	five := placeholder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}

	six := placeholder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}

	seven := placeholder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}

	eight := placeholder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}

	nine := placeholder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	separator := placeholder{
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	}

	digits := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	for line := range digits[0] {
		for digit := range digits {
			fmt.Print(digits[digit][line], "  ")
		}
		fmt.Println()
	}

	currentTime := time.Now()

	hours := currentTime.Hour()
	minutes := currentTime.Minute()
	seconds := currentTime.Second()

	hoursSplit := splitInt(hours)
	//fmt.Printf("%v", hoursSplit)
	minutesSplit := splitInt(minutes)
	//fmt.Printf("%v", minutesSplit)
	secondsSplit := splitInt(seconds)
	//fmt.Printf("%v", secondsSplit)

	var clock [8]placeholder
	indexFillInClock := 0
	//var hoursDigits [2]placeholder
	for _, value := range hoursSplit {
		clock[indexFillInClock] = digits[value]
		indexFillInClock++
	}

	clock[indexFillInClock] = separator
	indexFillInClock++

	//var minutesDigits [2]placeholder
	for _, value := range minutesSplit {
		clock[indexFillInClock] = digits[value]
		indexFillInClock++
	}

	clock[indexFillInClock] = separator
	indexFillInClock++

	//var secondsDigits [2]placeholder
	for _, value := range secondsSplit {
		clock[indexFillInClock] = digits[value]
		indexFillInClock++
	}

	for digitKey := range clock[0] {
		for elementKey := range clock {
			fmt.Printf("%s ", clock[elementKey][digitKey])
		}
		fmt.Println()
	}

}

func splitInt(n int) []int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n = n / 10
	}

	slcReverse := []int{}
	for i := len(slc) - 1; i >= 0; i-- {
		slcReverse = append(slcReverse, slc[i])
	}
	return slcReverse
}
