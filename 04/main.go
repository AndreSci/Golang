// Указатели. массивы и указатели. Цикл for. Массивы и слайсы

package main

import (
	"fmt"
	"errors"
)

var msg string

// Вызывается перед вызовом функции main
func init() {
	msg = "Hello I'm init"
}

func main() {
	fmt.Println(msg)
	fmt.Println("Hi")
	fmt.Println(errSome(1))

	message := "I will be Golang soon."
	printMessageCopy(message)
	fmt.Println(message)
	printMessage(&message)	// & - Refarance завернуть в казатель
	fmt.Println(message)

	writeData()

	initArray()
	matrixArray()
}

// Работа без указателя
func printMessageCopy(message string) {
	message += " Add new data copy"
	fmt.Println(message)
}
// Работа с указателем
func printMessage(message *string) {	// * - derefarance получить данные из указателя
	*message += " Add new data not copy"
	fmt.Println(*message)
}

func errSome(i int) (int, error) {
	fmt.Println(i)
	return 5, errors.New("WoW Wow guy")
}

func writeData() {
	var p* int
	number := 5

	p = &number

	fmt.Println(p)	// 0xc00000a108
	fmt.Println(*p)	// Данные 5
}

func initArray() {
	itemsInfo := [3]string{}	// На данном этапе 3 элемента массива будут заполнены нулем.
	fmt.Println(itemsInfo)

	items := [3]string{"1", "2", "3"}
	fmt.Println(items)

	// Слайсы
	// Слайсы это обёртка над массивом и ПРИ ПЕРЕДАЧИ в функцию мы работаем на прямую в базовом обьекте
	itemsSlice := []string{}
	itemsSlice = append(itemsSlice, "4")
	itemsSlice = append(itemsSlice, "2")
	fmt.Println(itemsSlice)

	itemsSliceMake := make([]string, 5)
	fmt.Println("Len: ", len(itemsSliceMake))
	fmt.Println("Cap: ", cap(itemsSliceMake))
}


func matrixArray() {
	matrix := make([][]int, 10)

	// цикл for основной в golang и может быть определёт условием как while так же как и for в C++ и так же как в Python
	for i := 0; i < 10; i++ {
		for k := 0; k < 10; k++{
			matrix[k] = make([]int, 10)	// создаем слайс с значением 0 в кажой ячейке
			matrix[k][i] = i
		}
		fmt.Println(matrix[i])
	}

	for i := range matrix {	// Получили index
		fmt.Println(matrix[i])
	}

	for index, value := range matrix {	// Получили index и value
		fmt.Println(index, " = ", value)
	}

	counter := 0
	// Цикл с прирыванием
	for {
		if counter == 10 {
			break
		}
		counter++
		fmt.Println(counter)
	}

	counter = 0
	// Цикл с прирыванием
	for counter < 10 {
		counter++
		fmt.Println(counter)
	}
}
