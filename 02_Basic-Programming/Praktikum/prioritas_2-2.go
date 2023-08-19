package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var input string
		fmt.Print("input number : ")
		fmt.Scanln(&input)

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}

		fmt.Printf("factor of %d : ", number)
		fmt.Println(" ")

		for i := 1; i <= number; i++ {
			if number%i == 0 {
				fmt.Println(i)
			}
		}
		break
	}

}
