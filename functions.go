package main

import (
	"fmt"
	"strings"
)

func PhoneticCurrent(array []int, size int) (res string) {
	var builder strings.Builder
	phoneticMap := PhoneticMap()
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
			stringToAdd := builder2.String()
			bigPhoneticMap[number] = stringToAdd
			fmt.Fprintf(&builder, "%s", stringToAdd)
		}
		fmt.Fprintf(&builder, "%s", ", ")
	}
	s := builder.String()   // no copying
	s = s[:builder.Len()-2] // no copying (removes trailing ", ")
	return s
}

func PhoneticPotentialImprovement(array []int, size int) (res string) {
	return PhoneticCurrent(array, size)
}

func PhoneticMap() (res map[int]string) {
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
	return phoneticMap
}
