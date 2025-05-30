package main

import (
	"fmt"
	"go-test/basic/cache"
)



func main() {
	fmt.Println("Hello")

	cache := cache.New()		// Почему cache.New() ? New это свободная функция которая в пакете создает экземпляр Структуры.....
	fmt.Println(cache)

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Set("userId2", 43)
	userId2 := cache.Get("userId2")

	fmt.Println(userId2)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)

	fmt.Println(cache)
}
