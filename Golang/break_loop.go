package main

import "fmt"

func main() {
	fmt.Println("Begin:")

LoopTag:
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break LoopTag
		}

		fmt.Println("Number is: ", i)
	}

	fmt.Println("End:")
}
