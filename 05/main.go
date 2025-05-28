package main

import "fmt"


func main() {
	// defer printMessage()	// defer метод предназначен для вызова в конце перед выходом из приложения
	defer handlerPanic()

	message := []string{"1", "2", "3"}

	fmt.Println(message)

	// panic("Help panic")

	printMessage()
}

func printMessage() {
	fmt.Println("pringMessage()")
}

// Нужен для обработки паники и продолжить работу приложения
func handlerPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
	fmt.Println("Done")
}
