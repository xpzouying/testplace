package main

import (
	"fmt"
	"os"
)

func main() {
	file1, err := os.Open("file1.txt")
	if err != nil {
		fmt.Println("Open file is: ", file1, "err!")
		return
	}
	defer file1.Close()
	buf := make([]byte, 100)
	for {
		n, _ := file1.Read(buf)
		if 0 == n {
			break
		}

		os.Stdout.Write(buf[:n])
	}
}
