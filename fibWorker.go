package main

import "fmt"

// ! w/ concurrency

func main() {
	num := 50
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i <= num; i++ {
		jobs <- i
	}
	close(jobs)

	fmt.Println("I HAVE WORKER GOROUTINES")
	for j := 0; j <= num; j++ {
		fmt.Println(j, " : ", <-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	// only send to jobs channel & receive on the results channel
	for n := range jobs {
		// send result of fib function to results channel
		results <- fib(n)
	}
}

func fib(n int) int {
	// fmt.Println("--", n)
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
