package main

import (
	"fmt"
	"time"
	"sync"
)

type counter struct {
	counter int
	mu *sync.Mutex	// так же есть RWMutex где доступны методы блокировки записи и свободное чтение RLock и RUnlock
}

func (c *counter) inc() {
	c.mu.Lock()
	c.counter ++
	c.mu.Unlock()
}

func (c* counter) value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}


// func main() {
// 	fmt.Println("Hi")

// 	counter := 0

// 	for i := 0; i <1000; i++ {
// 		go func() {
// 			counter++	// В этом смысле в один момент может несколько Goroutine 
// 			// попытаться переписать свой результат в переменную counter
// 		}()
// 	}

// 	time.Sleep(time.Second * 1)

// 	fmt.Println(counter)
// }

func main() {
	fmt.Println("Hi")

	counter := counter{ 
		mu: new(sync.Mutex),
	}

	for i := 0; i <1000; i++ {
		go func() {
			counter.inc()	// В этом смысле в один момент может несколько Goroutine 
			// попытаться переписать свой результат в переменную counter
		}()
	}

	time.Sleep(time.Second * 1)

	fmt.Println(counter.value())
}
