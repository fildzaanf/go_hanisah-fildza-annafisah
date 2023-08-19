package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		var (
			input string
			grade string
		)

		fmt.Print("input nilai : ")
		fmt.Scanln(&input)

		nilai, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}

		if nilai >= 80 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 65 && nilai <= 79 {
			grade = "B"
		} else if nilai >= 50 && nilai <= 64 {
			grade = "C"
		} else if nilai >= 35 && nilai <= 49 {
			grade = "D"
		} else if nilai >= 0 && nilai <= 34 {
			grade = "E"
		} else {
			grade = "nilai invalid"
		}
		fmt.Printf("output grade: %s", grade)
		break
	}

}
