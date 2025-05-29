package shape

import (
	"math"
)

// Неявный интерфейс	(В Go имплиментация не явная )
type Shape interface {
	Area() float32
}

// Словая структура и метод к ней
type Square struct {
	sideLenght float32
}

// Данная конструкция нужна для получения доступа к данным с нижним ригистром
func NewSquare(length float32) Square {
	return Square{
		sideLenght: length,
	}
}


func (s Square) Area() float32 {	// Создаем такую же функцию как в интерфейсе
	return s.sideLenght * s.sideLenght
}

// Словая структура и метод к ней
type Circle struct {
	radius float32
}

// Данная конструкция нужна для получения доступа к данным с нижним ригистром
func NewCircle(radius float32) Circle {
	return Circle{
		radius: radius,
	}
}

func (c Circle) Area() float32 {	// Создаем такую же функцию как в интерфейсе
	return c.radius * c.radius * math.Pi
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
