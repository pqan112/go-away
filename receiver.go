package main

import "fmt"

type Student struct {
	name  string
	age   int
	class string
}

func (s *Student) showInfo() {
	fmt.Printf("Ho ten: %s \n", s.name)
	fmt.Printf("Lop: %s \n", s.class)
	fmt.Printf("Ten: %d \n", s.age)
}

func (s *Student) clear() {
	s.age = 0
	s.class = ""
	s.name = ""
}

func main() {
	student1 := Student{
		name:  "An Pham",
		age:   10,
		class: "CNTT",
	}
	student1.showInfo()
	student1.clear()
	student1.showInfo()
}
