package main

import (
	"errors"
	"fmt"
)

/*
	для обработки ошибки используется пакет errors
	мы можем передать как ошибку errors.New() так и отсутсвие ошибки nil
	далее можно проверить если не ошибка то вывести сообщение, если ошибка вывести ошибку
*/

func main() {
	information, err := getPass(20)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(information)
}

func getPass(age int) (string, error) {
	if age >= 18 && age < 50 {
		return "Проходите", nil
	} else if age > 50 {
		return "Поберегите себя", errors.New("you are too old")
	}

	return "Вам еще нет 18", errors.New("you are too young")
}
