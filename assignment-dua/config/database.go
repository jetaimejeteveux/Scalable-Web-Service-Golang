package config

import (
	"assignment-dua/structs"
	"fmt"

	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	connection := "root:root123@tcp(127.0.0.1:3303)/assignmentdua?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connection)
	if err != nil {
		fmt.Println("Database Connection Failed")
		panic("Failed connect to database")
	}
	db.AutoMigrate(structs.Orders{})
	return db
}
