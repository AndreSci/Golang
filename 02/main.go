package main

import (
	"fmt"
)

func main() {
	printMessage("Message_1")
	printMessage("Message_2")
	printMessage("Message_3")

	fmt.Println(sayHello("Andrew"))
}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string) string {
	return "Hello, " + name		// Канкатенация строк
}
