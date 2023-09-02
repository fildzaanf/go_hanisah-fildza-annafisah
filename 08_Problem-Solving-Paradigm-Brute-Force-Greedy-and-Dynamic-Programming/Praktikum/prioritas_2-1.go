package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	cost := make([]int, len(jumps))
	cost[0] = 0
	cost[1] = int(math.Abs(float64(jumps[1] - jumps[0])))

	if len(jumps) <= 0 {
		return 0
	}

	for i := 2; i < len(jumps); i++ {
		cost[i] = int(math.Min(float64(cost[i-1]+int(math.Abs(float64(jumps[i]-jumps[i-1])))), float64(cost[i-2]+int(math.Abs(float64(jumps[i]-jumps[i-2]))))))
	}

	return cost[(len(jumps))-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}
