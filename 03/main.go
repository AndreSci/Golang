package main

import (
	"fmt"
	"errors"
)


func main() {

	fmt.Println("Hello")

	fmt.Println(prediction("Monday"))

	fmt.Println(switchPrediction("Tuesday"))


	fmt.Println(findMin(1,2,3,4,5,6,1,0))

	main2()

	main3()
}


func prediction(dayOfweek string) string {

	if dayOfweek == "Monday" {
		return "I hope this week will be good!"
	} else if dayOfweek == "Tuesday" {
		return "Second day of week"
	}

	return "Something went wrong"
}

func switchPrediction(datOfweek string) (string, error) {
	switch datOfweek {
	case "Monday":
		return "I hope this week will be good!", nil
	case "Tuesday":
		return "Second day of week", nil
	default:
		return "Something went wrong", errors.New("invalid day of week")
	}
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	} else {
		min := numbers[0]
		
		for _, i := range numbers {		// _, индекс в range // i это значение	 (очень удобно сразу иметь значение и индекс)
			fmt.Println(i)
			if min > i {
				min = i
			}
		}
		return min
	}
}

// АНОНИМНЫЕ ФУНКЦИИ --- ФУНКЦИЯ ЗАМЫКАНИЯ
func main2() {
	func() {
		fmt.Println("анонимная функция")
	} ()
}

// Круто экономит память 
func main3() {
	inc := incriment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}
// Такая функция запоминает состояние себя.
func incriment() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}
// Пока что сложно но прикольно
