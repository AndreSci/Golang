package main

import (
	"fmt"
	"test/basic/shape"
	"time"
)


func main() {
	// Тема интерфейсы
	// в Go все из пакета должно начинаться с большой буквы
	// 
	fmt.Println("Hi")
	time_item := time.Now()

	fmt.Println(time_item)
	square := shape.NewSquare(5)
	circle := shape.NewCircle(8)

	printShapeArea(square)
	printShapeArea(circle)

	// Пустому интерфейсу можно передать всё....
	printInterface(square)
	printInterface(44)
	printInterface("square")
	printInterface(true)
	// fmt.Println() так же принимает пустой интерфейс (по этой причине можно передавать любое колличество данных)

	// TYPE SWITCH
	printInterfaceTypeSwitch(square)
	printInterfaceTypeSwitch(44)
	printInterfaceTypeSwitch("square")
	printInterfaceTypeSwitch(true)

	// Привидения к типу
	printInterfaceTypeConversion("Hello")
	printInterfaceTypeConversion(11)
}


// Если передаваемая структура удовлетворяет условия интерфейса она выполнит медот данной структуры.
func printShapeArea(shape shape.Shape) {
	fmt.Println(shape.Area())
}

// ПУСТОЙ ИНТЕРФЕЙС

func printInterface(i interface{}) {
	fmt.Println(i)
}

func printInterfaceTypeSwitch(i interface{}) {
	switch value := i.(type) {
	case int:
		fmt.Println("int ", value)
	case bool:
		fmt.Println("bool ", value)
	default:
		fmt.Println("Unknown value")
	}
}


func printInterfaceTypeConversion(i interface{}) {
	value, ok := i.(string)

	if !ok{
		fmt.Println(value)
	} else {
		fmt.Println("Wrong type: interface is not string")
	}
}
