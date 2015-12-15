package main

import (
	"flag"
	"fmt"
	"net/url"
)

var (
	flagServer = flag.String("local", "fs://key:password@host", "File System")
)

func main() {
	flag.Parse()

	fmt.Println("flagServer: ", flagServer)

	u, err := url.Parse(*flagServer)
	if err != nil {
		fmt.Println("err!!!")
	}
	fmt.Println("url parse = ", u)
	fmt.Println("url.host= ", u.Host)

}
