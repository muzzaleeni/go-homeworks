package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	id int
}

type WorkerPool struct {
	numWorkers int
	jobs       chan Job
	results    chan int
	wg         sync.WaitGroup
}

func NewWorkerPool(n, m int) *WorkerPool {
	return &WorkerPool{
		numWorkers: n,
		jobs:       make(chan Job, m),
		results:    make(chan int, m),
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for job := range wp.jobs {
		// Do the actual work here
		fmt.Printf("Worker %d started job %d\n", id, job.id)
		time.Sleep(time.Second) // Simulating work
		fmt.Printf("Worker %d finished job %d\n", id, job.id)
		wp.results <- job.id
	}
}

func (wp *WorkerPool) AddJob(job Job) {
	wp.jobs <- job
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
	close(wp.results)
}

func (wp *WorkerPool) CollectResults() {
	for result := range wp.results {
		fmt.Printf("Job %d done\n", result)
	}
}

func main() {
	wp := NewWorkerPool(3, 10)

	for i := 0; i < 10; i++ {
		wp.AddJob(Job{id: i})
	}

	close(wp.jobs)

	wp.Start()
	wp.Wait()
	wp.CollectResults()
}
