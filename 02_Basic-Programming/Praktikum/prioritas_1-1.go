package main

import (
	"fmt"
	"strconv"
)

func calculateArea(base, top, height float64) float64 {
	return (0.5) * (base + top) * (height)
}

func main() {
	var (
		base, top, height float64
		input             string
		err               error
	)

	for {
		fmt.Print("input base   : ")
		fmt.Scanln(&input)
		base, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		break
	}

	for {
		fmt.Print("input top  : ")
		fmt.Scanln(&input)

		top, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		break
	}

	for {
		fmt.Print("input height   : ")
		fmt.Scanln(&input)

		height, err = strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		break
	}
	fmt.Println("Result       :", calculateArea(base, top, height))
}
