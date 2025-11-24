package main

import "fmt"

func main() {
	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
		"message 5",
	}


	// Альтернативный способ запуска цикла. Специально для слайсов и массивов
	// первый аргумент индекс, второй само значение
	for _, message := range messages {
		fmt.Println(message)
	}

	counter := 0

	// цикл без условия - бесконечный
	for {
		fmt.Println(counter)
		if(counter == 100) {
			break // <- ключевое слово для выхода из цикла
		}
		counter++
	}
}
