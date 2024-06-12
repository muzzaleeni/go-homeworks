package main

type WorkerPool struct {
	jobs chan int
	results chan int
	
}