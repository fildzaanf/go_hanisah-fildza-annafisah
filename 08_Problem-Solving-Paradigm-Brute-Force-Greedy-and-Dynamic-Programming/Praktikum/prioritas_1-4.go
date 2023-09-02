package main

import "fmt"

func SimpleEquations(a, b, c int) {
	max := 999
	for x := 1; x <= max; x++ {
		for y := 1; y <= max; y++ {
			for z := 1; z <= max; z++ {
				if (x+y+z == a) && (x*y*z == b) && (x*x+y*y+z*z == c) {
					fmt.Printf("%d %d %d\n", x, y, z)
					return
				}
			}
		}
	}

	fmt.Println("no solution")
}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
}
