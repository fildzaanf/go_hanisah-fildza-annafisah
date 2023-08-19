package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var input string
		fmt.Print("input rows : ")
		fmt.Scanln(&input)

		rows, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}

		fmt.Println("output     : ")
		fmt.Println(" ")

		for i := 1; i <= rows; i++ {
			for j := rows; j > i; j-- {
				fmt.Print(" ")
			}
			for k := 1; k <= i; k++ {
				fmt.Print("* ")
			}
			fmt.Println()
		}

		break

	}
}
