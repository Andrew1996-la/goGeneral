package main

import "fmt"

/*
	СЛАЙС
*/

func main() {
	mySlice := []string{"1", "2", "3"}

	fmt.Println(mySlice)

	changeArr(mySlice)

	fmt.Println(mySlice)
}

/*
	Слайс это обертка над массивом и он содержит ссылку.
	Поэтому передвая в функцию слайс и модифицируя его внутри
	фукнкции фактически мы модифицируем исходный слайс
*/
func changeArr(mySlice []string) {
	mySlice[1] = "13"
	fmt.Println(mySlice)
}
