package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	leftSum := 0
	rightSum := 0
	size := len(matrix)

	// diagonal left to right
	for i, row := range matrix {
		leftSum += row[i]
	}

	// diagonal right to left
	for i, row := range matrix {
		rightSum += row[size-1-i]
	}

	absoluteDifference := int(math.Abs(float64(leftSum - rightSum)))

	fmt.Printf("Matrix :\n")
	for _, row := range matrix {
		fmt.Println(row)
	}

	fmt.Printf("Diagonal Left to Right : %d\n", leftSum)
	fmt.Printf("Diagonal Right to Left : %d\n", rightSum)
	fmt.Printf("Absolute Difference    : %d\n", absoluteDifference)
}
