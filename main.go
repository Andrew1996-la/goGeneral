package main

import "fmt"

/*
СТРУКТУРЫ
*/

type User struct {
	id      int
	name    string
	married bool
}

func NewUser(id int, name string, married bool) User {
	return User{
		id:      id,
		name:    name,
		married: married,
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

	user2 := NewUser(2, "Ann", true) // <- оригинал
	user2.changeUser("Jula")         // <- модифицируем
	fmt.Println(user2)               // <- модифицированный user так как был передан pointer reciver
}
