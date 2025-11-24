package main

import "fmt"


/*
 defer откладывает выполнение функции до завершения других функций
 после чего начинает свое выполнение в порядке сверху вниз
*/

func main() {
	defer printText()
	fmt.Println("norm")
	defer printSomething()
	fmt.Println("text from main")
}

func printText()  {
	fmt.Println("any text from function")
}

func printSomething()  {
	fmt.Println("pring something")
}
