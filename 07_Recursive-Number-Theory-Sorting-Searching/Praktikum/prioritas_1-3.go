package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	if n == 2 || n == 3 || n == 5 {
		return true
	}

	if n%2 == 0 || n%3 == 0 || n%5 == 0 {
		return false
	}

	sqrt := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrt; i++ {
		if int(n)%i == 0 {
			return false
		}
	}
	return true
}

func primeX(number int) int {
	count := 0
	numb := 1

	if number <= 0 {
		return 0
	}

	for count < number {
		numb++

		if isPrime(numb) {
			count++
		}
	}
	return numb
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
