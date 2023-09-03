package main

import "fmt"

func primeX(number int) int {

	i := 2
	count := 0

	if number <= 0 {
		fmt.Print("invalid number : ")
		return number
	}

	for {

		isPrime := true

		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			count++
		}

		if count == number {
			return i
		}

		i++
	}
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
