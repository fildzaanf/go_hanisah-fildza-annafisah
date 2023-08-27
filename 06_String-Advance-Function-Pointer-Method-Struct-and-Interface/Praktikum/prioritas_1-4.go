package main

import "fmt"

func getMinMax(numbers ...*int) (min, max int) {
	min = *numbers[0]
	max = *numbers[0]

	for _, value := range numbers {

		if *value < min {
			min = *value
		}

		if *value > max {
			max = *value
		}
	}
	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int

	fmt.Println("Input angka : ")
	fmt.Scanln(&a1)
	fmt.Scanln(&a2)
	fmt.Scanln(&a3)
	fmt.Scanln(&a4)
	fmt.Scanln(&a5)
	fmt.Scanln(&a6)
	fmt.Println()

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai Min : ", min)
	fmt.Println("Nilai Max : ", max)
}
