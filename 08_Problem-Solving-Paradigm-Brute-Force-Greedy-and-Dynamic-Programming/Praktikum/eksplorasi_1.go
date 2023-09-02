package main

import (
	"fmt"
	"strconv"
)

func intToRoman(input int) string {

	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""

	for i := 0; i < len(romans); i++ {
		for input >= numbers[i] {
			result += romans[i]
			input -= numbers[i]
		}
	}

	return result
}

func main() {
	fmt.Println(intToRoman(4))    // IV
	fmt.Println(intToRoman(9))    // IX
	fmt.Println(intToRoman(23))   // XXIII
	fmt.Println(intToRoman(2021)) // MMXXI
	fmt.Println(intToRoman(1646)) // MDCXLVI
	for {
		var n string
		fmt.Print("input : ")
		fmt.Scanln(&n)

		input, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		output := intToRoman(input)
		fmt.Printf("output: %s\n\n", output)
		break
	}
}
