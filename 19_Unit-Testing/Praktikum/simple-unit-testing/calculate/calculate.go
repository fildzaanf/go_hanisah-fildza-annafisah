package calculate

import "errors"

type Calculate struct{}

func (c Calculate) Addition(number_a, number_b float64) float64 {
	result := number_a + number_b
	return result
}

func (c Calculate) Subtraction(number_a, number_b float64) float64 {
	result := number_a - number_b
	return result
}

func (c Calculate) Division(number_a, number_b float64) (float64, error) {
	if number_b == 0 {
		return 0, errors.New("can't be divided by zero")
	}
	result := number_a / number_b
	return result, nil
}

func (c Calculate) Multiplication(number_a, number_b float64) float64 {
	result := number_a * number_b
	return result
}
