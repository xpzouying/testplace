package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		fmt.Println("reader is: ", reader)
		return p[:n], nil
	}

	fmt.Println("No reader: ", p)
	return p, err
}

func main() {
	// fmt.Println("From stdin:")
	// data, err := ReadFrom(os.stdin, 11)
	// fmt.Println("DATA=", data)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("From file:")
	f, err := os.Open("./break_loop.go")
	data, err := ReadFrom(f, 9)
	if err != nil {
		panic(err)
	}
	fmt.Println("DATA=", data)

	fmt.Println("From file:")
	data, err = ReadFrom(strings.NewReader("from string"), 12)
	if err != nil {
		panic(err)
	}
	fmt.Println("DATA=", data)
}
