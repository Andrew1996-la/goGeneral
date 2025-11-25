package main

import "fmt"

/*
СТРУКТУРЫ
*/

/*
Часто структуры задаются именно таким образом
для удобства переиспользования
*/
type Car struct {
	manufacturer string
	cost         int
	color        string
	power        int
	hasLeftWeel  bool
}

/*
функция конструктор
*/
func NewCar(manufacturer, color string, cost, power int, hasLeftWeel bool) Car {
	return Car{
		manufacturer: manufacturer,
		color:        color,
		cost:         cost,
		power:        power,
		hasLeftWeel:  hasLeftWeel,
	}
}

func main() {

	/*
		один из способов инициализцации структуры.
		инициализируктся вместе со значением, не переиспользуется
	*/
	// car := struct {
	// 	manufacturer string
	// 	cost         int
	// 	color        string
	// 	power        int
	// 	hasLeftWeel  bool
	// }{
	// 	manufacturer: "BMW",
	// 	cost:         2000000,
	// 	color:        "Black",
	// 	power:        350,
	// 	hasLeftWeel:  true,
	// }

	/*
		инициализация структуры отнаследоавнной от Car
	*/
	car1 := Car{
		manufacturer: "Volvo",
		cost:         1500000,
		color:        "white",
		power:        180,
		hasLeftWeel:  true,
	}

	car2 := NewCar("Lada", "black", 250000, 80, true) // <- инициализация через конструктор
	
	

	fmt.Println(car1.cost) // <- обращение к конкретному свойству структуры по ключу
	fmt.Println(car1)
	fmt.Println(car2)

}
