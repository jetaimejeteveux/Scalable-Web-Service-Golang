package main

import (
	"encoding/json"
	"log"
	"net/url"
)

type User struct {
	FullName string `json:"name"`
	Age int

}

func main() {
	urlString := "https://google.com/search?keyword=golang&page=2"
	u, err := url.Parse(urlString)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Parsing URL ", u)
	log.Println("Scheme: ", u.Scheme)
	log.Println("hostname : ", u.Host)
	log.Println("path  : ", u.Path)
	log.Println("query string", u.Query()["page"])

	//encode-decode
	//encode from json to struct --> unmarshal
	//1. JSON TO -> body json || file -> string
	//2. Transform string -> byte
	//3. unmarshal to ..... map / interface  / struct

	jsonString :=  `{"name" : "ali", "age":12}`
	jsonData := []byte(jsonString)
	log.Println(string(jsonData))
	var userStruct User
	err = json.Unmarshal(jsonData, &userStruct)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(userStruct.FullName)

	// JSON to MAP (key:value)
	var userMap map[string]interface{}
	err = json.Unmarshal(jsonData, &userMap)
	log.Println(userMap["name"])

	var userIn interface{}
	err = json.Unmarshal(jsonData, &userIn)
	log.Println(userIn.(map[string]interface{})["name"])

	jsonStringArray := `[{"name":"ali", "age":12}, 
						{"name":"ali", "age":12}
						]`
	jsonStringArrayData := []byte(jsonStringArray)
	var userStructArray []User
	err = json.Unmarshal(jsonStringArrayData, &userStructArray)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(userStructArray[1].FullName)

	// object to string json
	var userObj = []User{
		{"ali", 23},
		{"lia", 24},
	}

	jsonDataByte, err := json.Marshal(userObj)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(string(jsonDataByte))
}
