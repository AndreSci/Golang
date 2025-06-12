package workerpool

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Job struct {
	URL string
}

type Result struct {
	URL          string
	StatusCode   int
	ResponseTime time.Duration
	Error        error
}

func (r Result) Info() string {
	if r.Error != nil {
		return fmt.Sprintf("[ERROR] - [%s] - %s", r.URL, r.Error.Error())
	}

	return fmt.Sprintf("[SUCCESS] - [%s] - Status: %d, REsponse Time: %s", r.URL, r.StatusCode, r.ResponseTime)
}

type Pool struct {
	worker      *worker
	workersCount int

	jobs   chan Job
	results chan Result

	wg      *sync.WaitGroup
	stopped bool
}

func New(workersCount int, timeout time.Duration, results chan Result) *Pool {
	return &Pool{
		worker:       newWorker(timeout),
		workersCount: workersCount,
		jobs:         make(chan Job),
		results:      results,
		wg:           new(sync.WaitGroup),
	}
}

func (p *Pool) Init() {
	for i := range p.workersCount {
		go p.initWorker(i)
	}
}

func (p *Pool) Push(j Job) {
	if p.stopped {
		return
	}

	p.jobs <- j
	p.wg.Add(1)
}
// Graceful Shutdown стиль
func (p *Pool) Stop() {
	p.stopped = true
	close(p.jobs)
	p.wg.Wait()
}
// Graceful Shutdown стиль
func (p *Pool) initWorker(id int) {
	for job := range p.jobs {
		time.Sleep(time.Second)
		p.results <- p.worker.process(job)
		p.wg.Done()
	}

	log.Printf("[worker ID %d] finished processing", id)
}
