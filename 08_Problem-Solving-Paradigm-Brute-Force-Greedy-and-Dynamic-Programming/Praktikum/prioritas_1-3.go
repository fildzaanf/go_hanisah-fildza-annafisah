package main

import "fmt"

var memory = make(map[int]int)

func fibonacci(number int) int {

	if value, found := memory[number]; found {
		return value
	}

	if number <= 1 {
		memory[number] = number
		return number
	}

	memory[number] = fibonacci(number-1) + fibonacci(number-2)
	return memory[number]
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}
