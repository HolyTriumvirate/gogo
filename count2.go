package main

import (
	"fmt"
	"time"
)

func main() {
	// go countTo(10)
	countTo(10)
	countTo(10)

	// fmt.Scanln()
}

func countTo(num int) {
	for i := 1; i <= num; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 250)
	}
}
