package main

import "fmt"

/*
СТРУКТУРЫ
*/

// Можем создавать кастомные типы чтобы для них сдлелать методы
type Age int 

// Метод для Age
func (a Age) isAdult() bool {
	return a >= 18
}

type User struct {
	id      int
	name    string
	married bool
	age     Age
}

func NewUser(id int, name string, married bool, age Age) User {
	return User{
		id:      id,
		name:    name,
		married: married,
		age: age,
	}
}

/*
Метод value reciver. Означает что метод принимает копию User и не может ее менять.
фактически выделяет новую ячейку памяти для User.
*/
func (u User) printUser(name string) {
	u.name = name
	fmt.Println(u.id, u.name, u.married)
}

/*
Метод pointer reciver. Означает что метод принимает указатель на оригинального User
и способен его модифицировать
*/
func (u *User) changeUser(name string) {
	u.name = name
	fmt.Println(u.id, u.name, u.married)
}

func main() {
	//user1 := NewUser(1, "John", false) // <- оригинал
	//user1.printUser("Peter")           // <- пробуем модифицировать
	//fmt.Println(user1)                 // <- остается все как было

	// user2 := NewUser(2, "Ann", true) // <- оригинал
	// user2.changeUser("Jula")         // <- модифицируем
	// fmt.Println(user2)               // <- модифицированный user так как был передан pointer reciver

	user3 := NewUser(3, "Serega", true, 21)
	fmt.Println(user3.age.isAdult())
}
