package main

import (
	"fmt"
	"sync"
)

var (
	factor = 1
	mutex  sync.Mutex
	wg     sync.WaitGroup
)

func factorial(number int) {
	output := 1
	defer wg.Done()

	for i := 1; i <= number; i++ {
		output *= i
	}

	mutex.Lock()
	factor *= output
	mutex.Unlock()
}

func main() {
	var num int

	fmt.Print("number: ")
	fmt.Scan(&num)

	wg.Add(num)

	for i := 1; i <= num; i++ {
		go factorial(i)
	}

	wg.Wait()

	fmt.Printf("Factorial result from %d is %d\n", num, factor)
}
