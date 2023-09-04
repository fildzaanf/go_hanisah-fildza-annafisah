package main

import (
	"fmt"
	"sync"
)

func multipleOfThree(number chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(number)
	max := 100

	for i := 1; i <= max; i++ {
		if i%3 == 0 {
			number <- i
		}
	}

}

func main() {

	number := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go multipleOfThree(number, &wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("multiple of 3 is : ")
		for value := range number {
			fmt.Println(value)
		}
	}()

	wg.Wait()
}
