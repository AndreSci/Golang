package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	printMessage("Message_1")
	printMessage("Message_2")
	printMessage("Message_3")

	fmt.Println(sayHello("Andrew"))

	var say_Demon string
	say_Demon = sayHello("Demon")
	fmt.Println(say_Demon)

	var say_Ira string
	say_Ira = sayHelloInt("Ira", 18)
	fmt.Println(say_Ira)

	var clubResp string
	var clubResult bool
	clubResp, clubResult = enterTheClub(15)

	fmt.Println(clubResp, " ", clubResult)

	// Без ошибки
	clubRespError, err := enterTheClubError(18)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(clubRespError)
	}
	// С ошибкой
	clubRespError, err = enterTheClubError(11)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(clubRespError)
	}
}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string) string {
	return "Hello, " + name		// Канкатенация строк
}


func sayHelloInt(name string, age int) string {
	// Обьединение строк и чисел в одну строку
	ret_value := fmt.Sprintf("Hello, %s! Your age is %d", name, age)
	return ret_value
}

func enterTheClub(age int) (string, bool) {	// Возврат нескольких значений
	var response string
	if age > 17 && age < 45 {
		response = "You can entry"
		return response, true
	} else if age >= 45 {
		return "You too old", false
	}

	response = "You can't go inside"
	return response, false
}


func enterTheClubError(age int) (string, error) {	// Возврат нескольких значений (с переменной Исключение ERROR)
	var response string
	if age > 17 && age < 45 {
		response = "You can entry"
		return response, nil
	} else if age >= 45 {
		return "You are too old", errors.New("you are too old for this")
	}
	
	response = "You can't go inside"
	return response, errors.New("you are too young for the club")
}
