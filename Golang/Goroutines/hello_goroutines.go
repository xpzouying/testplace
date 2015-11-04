package main

import (
	"fmt"
	"time"
)

func loop(n int) {
	fmt.Println("Goroutines: ", n)

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println("")
}

func main() {
	fmt.Println("Test1 - No goroutines")
	loop(1)
	loop(2)

	fmt.Println("Test2 - No time to call goroutines")
	go loop(3)
	loop(4)

	fmt.Println("Test3 - Have time to call goroutines")
	go loop(5)
	go loop(6)

	time.Sleep(time.Second)
}
