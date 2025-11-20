package main

import "fmt"

type Day int

const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	plan := getPlans(Sunday)

	fmt.Println(plan)
}

func getPlans(dayOfWeek Day) string {
	switch dayOfWeek {
	case Monday:
		return "Отличный день что бы начать то что давно откладывал"
	case Tuesday:
		return "Запланирую тренировку"
	case Wednesday:
		return "Подготовить доклад"
	case Thursday:
		return "Разобрать новую тему по go"
	case Friday:
		return "Немного отдохнем"
	case Saturday:
		return "Поездака за город"
	case Sunday:
		return "Запланируй следующюю неделю"
	default:
		return "не валидный день недели"
	}
}
