package main

import "fmt"

/*
	после того как функция завершается с паникой срабатывает handlePanic
	recover может опознать панику, и если паника есть мы не завершаем программу
	с паникой, а можем продолжить выполнение
*/
func main() {
	defer handlePanic()
	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
		"message 5",
	}

	messages[5] = "message 6"

	fmt.Println(messages)
}

func handlePanic() {

	/*
		if <инициализация>; <условие> {
	   		 // тело if
	}	*/
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
