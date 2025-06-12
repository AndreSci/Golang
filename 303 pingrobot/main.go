// #3 Pingrobot

package main

import (
	"chapter3/workerpool"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	INTERVAL        = time.Second * 10
	REQUEST_TIMEOUT = time.Second * 2
	WORKERS_COUNT   = 3
)

var urls = []string{
	"https://workshop.zhashkevych.com/",
	"https://academy.golang-ninja.com/",
	"https://zhaskevych.com/",
	"https://google.com/",
	"https://golang.org/",
}

func main() {

	results := make(chan workerpool.Result)
	workerPool := workerpool.New(WORKERS_COUNT, REQUEST_TIMEOUT, results)

	workerPool.Init()

	go generateJobs(workerPool)
	go processResults(results)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit // Чтение из канала это блокирующая операция

	workerPool.Stop()

}

func processResults(results chan workerpool.Result) {
	go func() {
		for result := range results {
			fmt.Println(result.Info())
		}
	}()
}

func generateJobs(wp *workerpool.Pool) {
	for {
		for _, url := range urls {
			wp.Push(workerpool.Job{URL: url})
		}

		time.Sleep(INTERVAL)
	}
}
