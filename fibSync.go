package main

import "fmt"

func main() {
	fmt.Println("I'M SINGLE THREADED")
	for i := 0; i <= 50; i++ {
		fmt.Println(i, " : ", fib(i))
	}
}

func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
