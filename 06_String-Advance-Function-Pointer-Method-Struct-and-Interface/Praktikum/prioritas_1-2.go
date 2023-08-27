package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	totalScore := 0
	for _, score := range s.score {
		totalScore += score
	}
	return float64(totalScore) / float64(len(s.score))
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	name = s.name[0]

	for i, score := range s.score {
		if score < min {
			min = score
			name = s.name[i]
		}
	}
	return min, name
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	name = s.name[0]

	for i, score := range s.score {
		if score > max {
			max = score
			name = s.name[i]
		}
	}
	return max, name
}

func main() {
	a := Student{}

	for i := 1; i < 6; i++ {

		var (
			name  string
			score int
		)

		fmt.Printf("Input %d Student's Name : ", i)
		fmt.Scan(&name)
		a.name = append(a.name, name)

		fmt.Printf("Input %d Student's Score : ", i)
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\nAverage Score : ", a.Average())

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score of Students : ", nameMin, "(", scoreMin, ")")
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score of Students : ", nameMax, "(", scoreMax, ")")
}
