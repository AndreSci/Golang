package main

import "fmt"


func main() {
	// ДЖЕНЕРИКИ 
	// два направления использования. Для функций и для обощенного момента
	fmt.Println("ДЖЕНЕРИКИ")
	fmt.Println("два направления использования. Для функций и для обощенного момента")

	a := []int64{1, 2,3}
	b := []float64{1.1, 2.2, 3.3}

	fmt.Println("--- We do not use Generic ---")
	fmt.Println(sumOfFloat64(b))
	fmt.Println(sumOfInt64(a))

	fmt.Println("--- We are using Generics now ---")
	fmt.Println(sum(a))
	fmt.Println(sum(b))

	fmt.Println("--- We are using Generics with interface now ---")
	fmt.Println(sumInterface(a))
	fmt.Println(sumInterface(b))

	fmt.Println("--- We are useing Interface/Generics now ---")
	fmt.Println(searchItem(a, 2))
	fmt.Println(searchItem(a, 55))

	d := []Users {
		{
			email: "some@gmail.com",
			name: "Nikita",
		},
				{
			email: "some2@gmail.com",
			name: "Andrew",
		},
				{
			email: "some3@gmail.com",
			name: "Den",
		},
	}
	fmt.Println("--- Type of test with User struct ---")
	fmt.Println(searchItem(d, Users{email: "some2@gmail.com",name: "Andrew",}))

	fmt.Println("--- Type of test with Any plus User struct ---")
	printAny(d)
}


func sumOfInt64(input []int64) int64 {
	var result int64

	for _, number := range input {
		result += number
	}

	return result
}


func sumOfFloat64(input []float64) float64 {
	var result float64

	for _, number := range input {
		result += number
	}

	return result
}

// ДЖЕНЕРИК

func sum[V int64 | float64](input []V) V {
	var result V

	for _, value := range input {
		result += value
	}

	return result
}

// ИНТЕРФЕЙС ДЛЯ GENERICS
type Number interface {
	int64 | float64
}

func sumInterface [V Number](input []V) V {
	var result V

	for _, value := range input {
		result += value
	}

	return result
}

// ВСТРОЕННЫЕ ИНТЕРФЕЙСЫ GENERECS
// ANY and COMPARABLE

// COMPARABLE
func searchItem [C comparable](elements []C, searchE C) bool {
	
	for _, value := range elements {
		if value == searchE {
			return true
		}
	}

	return false
}

// ANY

func printAny [T any](input T) {
	fmt.Println(input)
}


type Users struct {
	email string
	name string
}
