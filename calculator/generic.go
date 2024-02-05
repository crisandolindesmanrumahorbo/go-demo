package calculator

import "fmt"

type Number interface {
	int | float32
}

func Add[T Number](numbers ...T) T {
	// func Add[T int | float32](numbers ...T) T {
	var total T
	for _, val := range numbers {
		total += val
	}
	return total
}

func Print[T any, K any](first T, second K) string {
	result := fmt.Sprintf("%v is %v", first, second)
	return result
}
