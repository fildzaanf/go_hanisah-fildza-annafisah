package main

import (
	"fmt"
	"strconv"
)

func pascalTriangle(numRows int) [][]int {
	triangle := make([][]int, numRows)

	if numRows <= 0 {
		return [][]int{}
	}

	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}

func main() {
	for {
		var input string
		fmt.Print("input numRows : ")
		fmt.Scanln(&input)

		numRows, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}

		triangle := pascalTriangle(numRows)

		for _, row := range triangle {
			fmt.Println(row)
		}

		break
	}

}
