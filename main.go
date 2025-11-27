package main

import "fmt"

/*
Интерфейсы
*/

type Animal interface {
	MakeSound() string
	Move() string
}

type Lion struct {
	makeSound string
	move      string
}
type Monkey struct {
	makeSound string
	move      string
}
type Elephant struct {
	makeSound string
	move      string
}

func (l Lion) MakeSound() string {
	return "Lion make " + l.makeSound
}
func (l Lion) Move() string {
	return "Lion move " + l.move
}

func (m Monkey) MakeSound() string {
	return "Monkey make " + m.makeSound
}
func (m Monkey) Move() string {
	return "Monkey move " + m.move
}

func (e Elephant) MakeSound() string {
	return "Elephant make " + e.makeSound
}
func (e Elephant) Move() string {
	return "Elephant move " + e.move
}

func main() {
	animals := []Animal{
		Lion{makeSound: "rrr", move: "fast"},
		Monkey{makeSound: "bibibi", move: "middle"},
		Elephant{makeSound: "ooooo", move: "slow"},
	}

	for _, animal := range animals {
		printMove(animal)
	}
}
func printMove(animal Animal) {
	fmt.Println(animal.MakeSound())
	fmt.Println(animal.Move())
}

