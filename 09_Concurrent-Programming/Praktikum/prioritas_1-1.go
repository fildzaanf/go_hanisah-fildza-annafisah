package main

import (
	"fmt"
	"time"
)

func multipleX(x int, max int) {

	for i := 1; i <= max; i++ {
		fmt.Printf("multiple of %d times %d is %d\n", x, i, x*i)
		time.Sleep(3 * time.Second)
	}
}

func main() {

	var input, max int
	fmt.Print("input : ")
	fmt.Scanln(&input)
	fmt.Print("max   : ")
	fmt.Scanln(&max)

	go multipleX(input, max)
	time.Sleep(time.Duration(max) * 3 * time.Second)

}
