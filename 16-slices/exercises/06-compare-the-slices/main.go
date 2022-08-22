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
	"sort"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Compare the slices
//
//  1. Split the namesA string and get a slice
//
//  2. Sort all the slices
//
//  3. Compare whether the slices are equal or not
//
//
// EXPECTED OUTPUT
//
//   They are equal.
//
//
// HINTS
//
//   1. strings.Split function splits a string and
//      returns a string slice
//
//   2. Comparing slices: First check whether their length
//      are the same or not; only then compare them.
//
// ---------------------------------------------------------

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	nameSlice := strings.Split(namesA, ", ")
	//sort.Strings(nameSlice)
	sort.Sort(sort.Reverse(sort.StringSlice(nameSlice)))
	//sort.Strings(namesB)
	sort.Sort(sort.Reverse(sort.StringSlice(namesB)))

	if len(namesB) == len(nameSlice) {
		for i := range nameSlice {
			if nameSlice[i] != namesB[i] {
				fmt.Printf("%q\n", namesB[i])
				fmt.Printf("%q\n", nameSlice[i])
				fmt.Print("They are not equal")
				return
			}
		}
		fmt.Print("They are equal")
	}
}
