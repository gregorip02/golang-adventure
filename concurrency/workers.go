package main

import "fmt"

func main() {
	tasks := []int{2, 4, 5}
	workers := 3

	jobs := make(chan int, len(tasks))
	res := make(chan int, len(tasks))

	for i := 0; i < workers; i++ {
		go worker(i, jobs, res)
	}

	for _, value := range tasks {
		jobs<- value
	}

	close(jobs)

	for i := 0; i < len(tasks); i++ {
		<- res
	}
}

func worker(id int, jobs <-chan int, res chan<- int) {
	for job := range jobs {
		fmt.Printf("Woker with id %d started fib(%d)\n", id, job)
		fib := fibonacci(job)
		fmt.Printf("Worker with id %d was finished with fib(%d) = %d\n", id, job, fib)
		res<- fib
	}
}

func fibonacci(num int) int {
	if num <= 1 {
		return num
	}

	return fibonacci(num - 1) + fibonacci(num - 2)
}
