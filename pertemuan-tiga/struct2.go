package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var student1 = struct {
		grade int
		person
	}{grade: 100}

	student1.name = "ali"
	student1.age = 10
	fmt.Println(student1)

	//struct dalam slice

	var allPerson = []person{
		{name:"ali", age:100},
		{name:"lia", age:120},
	}
}
