package main

import "fmt"

var a, b, c int

func main() {
	information, entered := getPass(54)
	fmt.Println(information, entered)
}

func getPass(age int) (string, bool) {
	if age >= 18 && age < 50 {
		return "Проходите", true
	} else if age > 50 {
		return "Поберегите себя", false
	}

	return "Вам еще нет 18", false
}
