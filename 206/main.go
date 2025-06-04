package main

import (
	"context"
	"time"
	"fmt"
)

// 1. context.Background() - на самом высоком уровне. (желательно в функции main())
// 2. context.TODO - когда не уверены, какой контекст использовать. (заменить в проде)
// 3. context.Value - стоит использовать как можно реже и передавать только необязательные параметры.
// 4. context всегда передается первым аргументом в функцию.

func main() {
	// Background это шаблон для создания контекста, ничего не содержит не является nil
	// Создается в функциях main() или верхнем уровне задачи

	ctx := context.Background()
	// Контекст может принимать результат предыдущего контекста как матрёшка

	// go func () {
	// 	time.Sleep(time.Millisecond*100)
	// 	cancel()
	// }()


	parse(ctx)
}

func parse(ctx context.Context) {

	ctx, _ = context.WithTimeout(ctx, time.Second * 3)
	ctx = context.WithValue(ctx, "id", 1555)

	id := ctx.Value("id")
	fmt.Println("id:", id)

	// res := 5 + id.(int)
	// fmt.Println(res)
	
	for {
		select {
		case <- time.After(time.Second*2):
			fmt.Println("Parsing completed")
			return
		case <-ctx.Done():
			fmt.Println("deadline exceded")
			return
		}
	}
}
