package main

import (
	"fmt"
	"sync"
)

func multipleOfThree(number chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	max := 100

	for i := 1; i <= max; i++ {
		if i%3 == 0 {
			number <- i
		}
	}
}

func main() {

	var wg sync.WaitGroup
	number := make(chan int)

	wg.Add(1)
	go multipleOfThree(number, &wg)

	go func() {
		defer close(number)
		fmt.Println("multiple of 3 is : ")
		for value := range number {
			fmt.Println(value)
		}
	}()

	wg.Wait()
}
