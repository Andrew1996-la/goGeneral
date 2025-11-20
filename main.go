package main

import "fmt"

/*
	ЗАМЫКАНИЕ

	В примере ниже реализован пример замыкания

	функция interment возвращает функцию которая возвращает int
	внутри interment есть переменная count которая сохраняется и инкрементируется при вызове ананимной функции

	counter := interment(). в counter записывается функция
			return func() int {
			count++
			return count
		}
	по сути это ананимная функция из interment

	каждый вызов функции counter инкрементирует counter.

*/

func main() {
	counter := interment()

	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func interment() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}
