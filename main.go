package main

import "fmt"

/*
	МАПА
*/
func main() {

	/*
		первый способ инициализации мапы
		создаеься и можем сразу положить данные
	*/
	users:= map[string]int{
		"Joe": 29,
		"Doy": 27,
		"Peter": 35,
	}
	
	delete(users, "Joe") // <- Удаление из мапы
	users["Ann"] = 23 // <- Добавление нового значения

	/*
		из мапы мы можем записать переменную. По ключу получаем значение
		exists или ok это признак что ключ был наиден. Если ключа нет - false
	*/
	age, exists := users["Peterа"]

	if(!exists) {
		fmt.Println("Такого пользователя нет")
	} else {
		fmt.Println(age)
	}


	/*
		цикл для перебора мапы
		можно деструктуризировать ключ и значение
	*/
	for name, age := range users {
		fmt.Println(name, age)
	}


	/* 
		второй способ инициализации мапы (make)
		при инициализации данные положить не можем
	*/
	// cars := make(map[string]int)


}

