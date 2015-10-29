package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse args
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("---Debug: ", k, v)

		fmt.Println("Key-: ", k)
		fmt.Println("Val-: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello Web Server! power by zy")
}

func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
