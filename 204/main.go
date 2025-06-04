package main

import (
	"fmt"
	"time"
)

func main() {
	// создаем канал
	var msg chan string

	// msg <- "Test writting"   // Ошибка deadlock (без горутины невозможно вписать данные в канал)

	msg = make(chan string, 3)	// Добавления размера буффера снимает ситуацию DEADLOCK

	// Анонимная функция
	go func() {
		time.Sleep(time.Second)
		msg <- "Hello 1"
		msg <- "Hello 2"
		msg <- "Hello 3"
	}()
	
	// value := <- msg
	// fmt.Println(value)

	fmt.Println(<-msg)
	fmt.Println(<-msg)
	fmt.Println(<-msg)

	// fmt.Println(<-msg)  4й вызов сделает DEADLOCK программмы и будет ждать появления данных в канале

	forChannel()

	forChannelEasyStyle()

	workWithTwoChannals()
}


func forChannel() {

	msg := make(chan string, 3)

	msg <- "Func hello 1"
	msg <- "Func hello 2"
	msg <- "Func hello 3"

	close(msg)

	for {
		value, ok := <- msg		// Для такого метода нужно закрывать канал
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println(value)
	}
}

func forChannelEasyStyle() {
	// Вариант с обычными циклом 
	msg := make(chan string, 3)

	msg <- "Func Easy hello 1"
	msg <- "Func Easy hello 2"
	msg <- "Func Easy hello 3"

	close(msg)

	for m := range msg {
		fmt.Println(m)
	}
}


func workWithTwoChannals() {
	message1 := make(chan string)
	message2 := make(chan string)

	go func() {
		for {
			message1 <- "Channal 1. Went 200 ms."
			time.Sleep(time.Millisecond* 200)
		}
	}()

		go func() {
		for {
			message2 <- "Channal 2. Went 1 sec."
			time.Sleep(time.Second)
		}
	}()

	// for {
	// 	// Операция <- блокирующая
	// 	fmt.Println(<- message1)
	// 	fmt.Println(<- message2)
	// }

	for {
		// метод select не блокирует вывод из канала
		select {
		case msg := <-message1:
			fmt.Println(msg)
		case msg := <-message2:
			fmt.Println(msg)
		}
	}
}
