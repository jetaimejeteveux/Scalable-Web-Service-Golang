package main

import "fmt"

type person struct {
	name string
	age  int
}
type student struct {
	name  string
	grade int
	person
}

func main() {
	var studentA student

	studentA.name = "ali"
	studentA.grade = 100

	fmt.Println(studentA)
	fmt.Println(studentA.name)
	fmt.Println(studentA.grade)

	var personB = person{name: "ila", age: 12}
	var studentD = student{person: personB, grade: 100}
	var studentB = student{name: "lia", grade: 90}

	fmt.Println(studentB)
	fmt.Println(studentB.name)
	fmt.Println(studentB.grade)

	var studentC *student = &studentB
	studentC.name = "namaC"
	studentC.grade = 80

	fmt.Println(studentC)
	fmt.Println(studentC.name)
	fmt.Println(studentC.grade)
	// b ikut berubah
	fmt.Println(studentB)
	fmt.Println(studentB.name)
	fmt.Println(studentB.grade)

	fmt.Println(studentD)
	fmt.Println(studentD.name)
	fmt.Println(studentD.grade)

	
}
