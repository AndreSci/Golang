package main

import "fmt"


var glob_b int

func main() {

	var item string

	item = "Andrew"
	fmt.Println("Hi", item)


	// Раздел выделения памяти в куче
	a := createPtr()

	fmt.Println(*a)

	// Наглядный пример изменения данных в куче через указатель
	*a = 10
	fmt.Println(*a)
}

func createPtr() *int {

	a := 5
	return &a
}
