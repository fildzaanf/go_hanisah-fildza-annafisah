package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var input string
		fmt.Print("input bilangan : ")
		fmt.Scanln(&input)

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}

		fmt.Printf("Bilangan %d adalah ", number)

		if number%2 == 0 {
			fmt.Println("Bilangan Genap")
		} else {
			fmt.Println("Bilangan Ganjil")
		}

		break
	}
}
