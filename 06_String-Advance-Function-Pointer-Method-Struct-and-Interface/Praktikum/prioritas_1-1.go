package main

import "fmt"

type Car struct {
	carType string
	fuelin  float64
}

func (car Car) estimateDistance() float64 {
	return car.fuelin * (1.5)
}

func main() {
	myCar := Car{"SUV", 10.5}
	myEstimateDistance := myCar.estimateDistance()

	fmt.Printf("Car type : %s, Estimate distance : %.2f\n", myCar.carType, myEstimateDistance)
}
