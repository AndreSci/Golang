package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Первое обьявление string через :=
	msg := "Я скоро стану go разрабом"
	fmt.Println(msg)
	// Прямое определение типа переменной
	var item string = "Hi"
	fmt.Println(item)
	// Константрая переменная
	const item2 string = "Hello const"
	fmt.Println(item2)
	// Числовой тип
	const item3 int = 2
	fmt.Println(item3)
	// Инициализация переменной без данных
	var item4 string
	item4 = "Some string"
	fmt.Println(item4)
	// Метод как узнать тип переменной
	fmt.Println(reflect.TypeOf(item4))
	// float32
	var f_numb float32
	f_numb = 3.11
	fmt.Println(f_numb)

	// BOOL
	var b_item bool
	b_item = true
	fmt.Println(b_item)

	// BYTES	хранение в СЛАЙСАХ
	bytes_mass := []byte("some data")
	fmt.Println(bytes_mass)
	// BYTE
	var byte_item byte = 97
	fmt.Printf("%c\n", byte_item)
	// RUNE = int32 (редко и сложно)
	var r_item rune = 'a'
	fmt.Println(r_item)

	// Массовая инициализация
	a, b, c := 1, 2, 3
	fmt.Println(a, b, c)
	// Тут же перестановка местами без потерь в памяти
	b, c, a = a, b, c
	fmt.Println(a, b, c)
	// отправляем в пустату значение
	a, _, c = 1, 2, 3
	fmt.Println(a, b, c)
}
