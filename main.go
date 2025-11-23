package main

import "fmt"

func main() {
	//var mySlice []string // <- при такой инициализации слайса его указатель nil
	mySlice := make([]string, 50)

	mySlice = append(mySlice, "51")

	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
}
