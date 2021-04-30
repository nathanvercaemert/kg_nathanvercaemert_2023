package main

import (
	"fmt"
	"strings"
)

func PhoneticCurrent(array []int, size int) (res string) {
	var builder strings.Builder

	phoneticMap := make(map[int]string)
	phoneticMap[0] = "Zero"
	phoneticMap[1] = "One"
	phoneticMap[2] = "Two"
	phoneticMap[3] = "Three"
	phoneticMap[4] = "Four"
	phoneticMap[5] = "Five"
	phoneticMap[6] = "Six"
	phoneticMap[7] = "Seven"
	phoneticMap[8] = "Eight"
	phoneticMap[9] = "Nine"

	bigPhoneticMap := make(map[int]string)
	bigPhoneticMap[0] = "Zero"

	for i := 0; i < size; i++ {
		var number int = array[i]
		if val, ok := bigPhoneticMap[number]; ok {
			fmt.Fprintf(&builder, "%s", val)
		} else {
			var builder2 strings.Builder
			var builder3 strings.Builder
			sameNumber := number
			for sameNumber > 0 {
				modulo := sameNumber % 10
				sameNumber /= 10
				builder3.Reset()
				fmt.Fprintf(&builder3, "%s", phoneticMap[modulo])
				fmt.Fprintf(&builder3, "%s", builder2.String())
				builder2 = builder3
			}
			bigPhoneticMap[number] = builder2.String()
			fmt.Fprintf(&builder, "%s", builder2.String())
		}
		fmt.Fprintf(&builder, "%s", ", ")
	}
	// s := builder.String()   // no copying
	// s = s[:builder.Len()-2] // no copying (removes trailing ", ")
	// return s
	return builder.String()[:builder.Len()-2]
}

func PhoneticPotentialImprovement(array []int, size int) (res string) {
	var builder strings.Builder

	bigPhoneticMap := make(map[int]string)
	bigPhoneticMap[0] = "Zero"

	for i := 0; i < size; i++ {
		var number int = array[i]
		if val, ok := bigPhoneticMap[number]; ok {
			fmt.Fprintf(&builder, "%s", val)
		} else {
			var builder2 strings.Builder
			var builder3 strings.Builder
			sameNumber := number
			for sameNumber > 0 {
				modulo := sameNumber % 10
				sameNumber /= 10
				builder3.Reset()
				switch modulo {
				case 1:
					fmt.Fprintf(&builder3, "%s", "One")
				case 2:
					fmt.Fprintf(&builder3, "%s", "Two")
				case 3:
					fmt.Fprintf(&builder3, "%s", "Three")
				case 4:
					fmt.Fprintf(&builder3, "%s", "Four")
				case 5:
					fmt.Fprintf(&builder3, "%s", "Five")
				case 6:
					fmt.Fprintf(&builder3, "%s", "Six")
				case 7:
					fmt.Fprintf(&builder3, "%s", "Seven")
				case 8:
					fmt.Fprintf(&builder3, "%s", "Eight")
				case 9:
					fmt.Fprintf(&builder3, "%s", "Nine")
				case 0:
					fmt.Fprintf(&builder3, "%s", "Zero")
				}
				fmt.Fprintf(&builder3, "%s", builder2.String())
				builder2 = builder3
			}
			bigPhoneticMap[number] = builder2.String()
			fmt.Fprintf(&builder, "%s", builder2.String())
		}
		fmt.Fprintf(&builder, "%s", ", ")
	}
	// s := builder.String()   // no copying
	// s = s[:builder.Len()-2] // no copying (removes trailing ", ")
	// return s
	return builder.String()[:builder.Len()-2]
}
