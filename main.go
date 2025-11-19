package main

import "fmt"

var a, b, c int

func main() {
	userInfo := getUserInfo("Ivan", 40)
	printMessage(userInfo)
}

func printMessage(message string) {
	fmt.Println(message)
}

func getUserInfo(name string, age int) string {
	return fmt.Sprintf("Hello! I am %s! And I am %d years old", name, age)
}