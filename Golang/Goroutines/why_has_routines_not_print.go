package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
		runtime.Gosched()
	}
}

func why_goroutines_not_print() {
	fmt.Println("-----Begin why_goroutines_not_print():")
	defer fmt.Println("-----End why_goroutines_not_print():")

	// Not print out
	go say("Hello")
	go say("World")
}

func main() {
	fmt.Printf("This CPU has %d cores.\n", runtime.NumCPU)
	// runtime.GOMAXPROCS(2)

	why_goroutines_not_print()

	time.Sleep(1000 * time.Millisecond)
}
