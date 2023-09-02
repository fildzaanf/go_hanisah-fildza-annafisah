package main

import (
	"fmt"
	"strconv"
)

func binaryNumber(n int) []string {
	ans := make([]string, n+1)

	for i := 0; i <= n; i++ {
		ans[i] = strconv.FormatInt(int64(i), 2)
	}
	return ans
}

func main() {
	fmt.Println(binaryNumber(2)) // [0 1 10]
	fmt.Println(binaryNumber(3)) // [0 1 10 11]
	for {
		var input string
		fmt.Print("input n : ")
		fmt.Scanln(&input)

		n, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		output := binaryNumber(n)
		fmt.Print(output)
		break
	}
}
