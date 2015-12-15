package main

import (
	"fmt"
)

type Student struct {
	name string
}

func (s *Student) ChangeName() {
	// s.name = "xxx"
	(*s).name = "xxx"
}

func main() {
	stu1 := Student{"stu1"}
	stu1.ChangeName()
	fmt.Println("ChangedName = ", stu1.name)
}
