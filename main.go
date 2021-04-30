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

func Phonetic(slice []string, size int) (res string) {
	var builder strings.Builder

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

	bigPhoneticMap := make(map[string]string)
	bigPhoneticMap["0"] = "Zero"

	for i := 1; i < size; i++ {
		var number string = slice[i]
		if val, ok := bigPhoneticMap[number]; ok {
			fmt.Fprintf(&builder, "%s", val)
		} else {
			var builder2 strings.Builder
			var builder3 strings.Builder
			sameNumber, _ := strconv.Atoi(number)
			for sameNumber > 0 {
				modulo := sameNumber % 10
				sameNumber /= 10
				builder3.Reset()
				fmt.Fprintf(&builder3, "%s", phoneticMap[strconv.Itoa(modulo)])
				fmt.Fprintf(&builder3, "%s", builder2.String())
				builder2 = builder3
			}
			bigPhoneticMap[number] = builder2.String()
			fmt.Fprintf(&builder, "%s", builder2.String())
		}
		fmt.Fprintf(&builder, "%s", ", ")
	}
	return builder.String()[:builder.Len()-2]
}
