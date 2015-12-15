package main

import (
	"fmt"
	"reflect"
)

type UserWithTag struct {
	name    string            "User Name"
	age     int               "User Age"
	address map[string]string "City --> Info"
}

func reflectTag(u UserWithTag, i int) {
	uType := reflect.TypeOf(u)
	uField := uType.Field(i)

	fmt.Println("User type:", uType, " field:", uField)

}

//////////////////////////////
// Struct
type body struct {
	name       string
	age        int
	girlfriend string
}

type family struct {
	string "family name"
	body   "the embedded structs"
}

func (f family) changeFamilyNameByVar(s string) {
	f.string = "This Family Name has changed."
}

func (f *family) changeFamilyNameByPointer(s string) {
	f.string = "This Family Name has changed."
}

func testVarOrPointerName() {
	fmt.Println("=====By value=====")
	f1 := family{"Noname", body{}}
	fmt.Println("Before changed:", f1)
	f1.changeFamilyNameByVar("ChangedName")
	fmt.Println("After changed:", f1)

	fmt.Println("=====By pointer=====")
	f2 := new(family)
	f2.string = "NONAME"
	fmt.Println("Before changed:", f2)
	f2.changeFamilyNameByPointer("ChangedName")
	fmt.Println("After changed:", f2)
}

func testAnonymousFunc() {
	fmt.Println("testAnonymousFunc(): ")

	var f *family = new(family)
	f.string = "One World"
	f.name = "Im host"
	f.age = 30
	f.girlfriend = "left"

	fmt.Println("here is a family: ", f)

	f2 := family{"Two World", body{"body", 99, "girl2"}}
	fmt.Println("here is a family: ", f2)

	// echo literal expression & print anonymous field
	f2Type := reflect.TypeOf(f2)
	fieldNumber := f2Type.NumField()
	for i := 0; i < fieldNumber; i++ {
		fmt.Println("field index = ", i,
			", field content: ", f2Type.Field(i),
			", field tag: ", f2Type.Field(i).Tag)
	}

}

//////////////////////////////

func main() {
	var user1 *UserWithTag = new(UserWithTag)
	user1.name = "user1"
	user1.age = 18
	user1.address = map[string]string{"Beijing": "Beijing info",
		"Shanghai": "Shanghai Info"}
	// error:
	// user1.address["beijing"] = "Beijing Address info"
	// user1.address["shanghai"] = "Shanghai Address info"

	fmt.Println("user1 info: ", user1)

	testAnonymousFunc()

	testVarOrPointerName()
}
