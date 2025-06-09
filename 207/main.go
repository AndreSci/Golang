package main

import (
	"fmt"
	"time"
)


func main() {
  // Тема WORKER POOL
	t := time.Now()
	const jobsCount, workerCount = 15, 15

	jobs := make(chan int, 15)

	results := make(chan int, 15)

	for i := range workerCount {
		go worker(i + 1, jobs, results)
	}

	for i := range jobsCount {
		jobs <- i +1
	}

	close(jobs)

	for i := range jobsCount {
		fmt.Printf("result #%d : value = %d\n", i+ 1, <-results)
	}

	
	fmt.Println("TIME ELAPSED:", time.Since(t).String())
}
// Сложная тема с стрелочками в каналах где указания стрелочки определяет уровень доступа(чтение\запись)
func worker(id int, jobs <- chan int, results chan<- int) {
	for j := range jobs {
		time.Sleep(time.Second)
		fmt.Printf("worker #%d finished\n", id)
		results <- j * j
	}
}
