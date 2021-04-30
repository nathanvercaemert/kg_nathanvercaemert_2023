package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print(Phonetic(os.Args, len(os.Args)))
}

// Determined to be the most effective implementation.
func Phonetic(slice []string, size int) (res string) {
	var builder strings.Builder

	// Map of each digit to its phonetic equivalent.
	phoneticMap := make(map[string]string)
	phoneticMap["0"] = "Zero"
	phoneticMap["1"] = "One"
	phoneticMap["2"] = "Two"
	phoneticMap["3"] = "Three"
	phoneticMap["4"] = "Four"
	phoneticMap["5"] = "Five"
	phoneticMap["6"] = "Six"
	phoneticMap["7"] = "Seven"
	phoneticMap["8"] = "Eight"
	phoneticMap["9"] = "Nine"

	// Map of each integer to its phonetic equivalent.
	bigPhoneticMap := make(map[string]string)
	// Zero is included beforehand to assist with later logic.
	bigPhoneticMap["0"] = "Zero"

	// For each of the arguments.
	for i := 1; i < size; i++ {
		var number string = slice[i]
		// If the argument is already in the map of integers:
		if val, ok := bigPhoneticMap[number]; ok {
			// Get the phonetic equivalent for the value.
			fmt.Fprintf(&builder, "%s", val)
		} else {
			var builder2 strings.Builder
			var builder3 strings.Builder
			sameNumber, _ := strconv.Atoi(number)
			// For each digit (starting with smallest:)
			for sameNumber > 0 {
				// Determine the digit.
				modulo := sameNumber % 10
				// Move to the next digit for the next iteration.
				sameNumber /= 10
				builder3.Reset()
				// Prepend the phonetic equivalent of the digit.
				fmt.Fprintf(&builder3, "%s", phoneticMap[strconv.Itoa(modulo)])
				fmt.Fprintf(&builder3, "%s", builder2.String())
				builder2 = builder3
			}
			// Add the phonetic equivalent of the argument to the map.
			bigPhoneticMap[number] = builder2.String()
			// Add the phonetic equivalent to the output.
			fmt.Fprintf(&builder, "%s", builder2.String())
		}
		fmt.Fprintf(&builder, "%s", ", ")
	}
	// Don't forget to remove the last ", "
	return builder.String()[:builder.Len()-2]
}
