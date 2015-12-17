package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	b := []byte(`{"Name": "zouying", "Age": 30, "Friends": ["buddy1", "buddy2"]}`)

	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("Json format error")
		panic(err)
	}

	fmt.Println("interface type==>", reflect.TypeOf(f))
	fmt.Println("interface==>", f)

	m := f.(map[string]interface{})

	for k, v := range m {
		fmt.Println("---------")
		fmt.Println("In interface: key=", k)
		fmt.Println("In interface: value=", v)

		switch vv := v.(type) {
		case string:
			fmt.Println(k, ",is string, ", vv)
		case int:
			fmt.Println(k, ",is int, ", vv)
		case []interface{}:
			fmt.Println(k, ",is array, ", vv)
		}
	}
}
