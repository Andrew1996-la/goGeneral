package main

import "fmt"

func main() {
	fmt.Println(getMin(3, 4, 2, 5, 6, 74, 76, -2))
}

func getMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}


	return min
}
