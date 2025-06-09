package main

import (
	"fmt"
	"go-test/basic/cache"
	"log"
	"time"
)



func main() {
	fmt.Println("Hello")

	cache := cache.New()		// Почему cache.New() ? New это свободная функция которая в пакете создает экземпляр Структуры.....
	fmt.Println(cache)

	cache.Set("userId", 42, time.Second * 5)

	userId, err := cache.Get("userId")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-- First Get:", userId)
	time.Sleep(time.Second * 2)

	userId2, err2 := cache.Get("userId")

	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("-- Second Get:", userId2)

	time.Sleep(time.Second * 4)

	userId3, err3 := cache.Get("userId")

	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println("-- Third Get:", userId3)
}
