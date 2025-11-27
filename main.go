package main

import (
	"fmt"
	"math"
)

/*
Композиция интерфейсов
*/

/*
Интерфейс содержащий другие интерфейсы называется композиция интерфейсов

Когда функция принимает аргументом shape то ее структуры должны иметь
имеплементация и ShapeWithArea и ShapeWithPrimetr
*/
type Shape interface {
	ShapeWithArea
	ShapeWithPrimetr
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPrimetr interface {
	Perimetr() float32
}

type Square struct {
	sideLength float32
}

type Circle struct {
	radius float32
}

/*
для Square и Circle имплементируем метод Area
*/
func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}
func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

/*
для Square и Circle имплементируем метод Perimetr
*/
func (s Square) Perimetr() float32 {
	return s.sideLength * 4
}
func (c Circle) Perimetr() float32 {
	return 2 * c.radius * c.radius * math.Pi
}

func main() {
	square := Square{sideLength: 5}
	circle := Circle{radius: 7}

	printShapeArea(square)
	printShapeArea(circle)
}


func printShapeArea(s Shape) {
	fmt.Println(s.Area())
	fmt.Println(s.Perimetr())
}
