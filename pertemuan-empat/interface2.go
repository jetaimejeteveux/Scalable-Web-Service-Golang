package main

import "fmt"

type person struct {
	name  string
	grade int
}

func main() {
	var unknownType interface{}

	unknownType = "hellow world"
	unknownType = 100

	data := map[string]interface{}{
		"nama":   "ali",
		"bahasa": "go",
		"grade":  100,
	}
	fmt.Println(data)

	var unknownStruct interface{} = &person{name: "ali", grade: 10}
	//fmt.Println(unknownStruct.name) umumnya begini tapi error
	fmt.Println(unknownStruct.(*person).name)

	var unkownSlice interface{} = []string{"a", "b", "c"}
	fmt.Println(unkownSlice.([]string)[1])
}
