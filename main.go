package main

import "fmt"

/*
Интерфейсы
*/
type Calculator interface {
	Calculate(num1, num2 int) int
	Name() string
}

type Adder struct{}
type Multiplier struct{}
type Subtractor struct{}

func (a Adder) Calculate(num1, num2 int) int {
	return num1 + num2
}
func (a Adder) Name() string {
	return "Adder"
}

func (m Multiplier) Calculate(num1, num2 int) int {
	return num1 * num2
}
func (m Multiplier) Name() string {
	return "Multiplier"
}

func (s Subtractor) Calculate(num1, num2 int) int {
	return num1 - num2
}
func (s Subtractor) Name() string {
	return "Subtractor"
}

func main() {
	calculator := []Calculator{
		Adder{},
		Multiplier{},
		Subtractor{},
	}

	a, b  := 10, 6
	
	for _, calculate := range calculator {
		result := calculate.Calculate(a, b)

        fmt.Printf("%s: %d и %d = %d\n", calculate.Name(), a, b, result)
	}

}


