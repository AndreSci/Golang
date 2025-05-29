package main

import (
	"fmt"
	"math"
)


func main() {
	// Тема интерфейсы

	fmt.Println("Hi")

	square := Square{5}
	circle := Circle{8}

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

// Неявный интерфейс	(В Go имплиментация не явная )
type Shape interface {
	Area() float32
}

// Словая структура и метод к ней
type Square struct {
	sideLenght float32
}
func (s Square) Area() float32 {	// Создаем такую же функцию как в интерфейсе
	return s.sideLenght * s.sideLenght
}

// Словая структура и метод к ней
type Circle struct {
	radius float32
}
func (c Circle) Area() float32 {	// Создаем такую же функцию как в интерфейсе
	return c.radius * c.radius * math.Pi
}

// Если передаваемая структура удовлетворяет условия интерфейса она выполнит медот данной структуры.
func printShapeArea(shape Shape) {
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


// Множественные интерфейсы	(композиция интерфейсов)	Тип должен содержать все методы всех интерфейсов
// На данном случаи кастомный тип должен иметь Area2() и Perimeter() методы...
type ShapeTwo interface {
	ShapeWithArea2
	ShapeWithPerimeter
}

type ShapeWithArea2 interface {
	Area2()
}

type ShapeWithPerimeter interface {
	Perimeter()
}
