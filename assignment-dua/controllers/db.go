package controllers

import "github.com/jinzhu/gorm"

type DBconn struct {
	DB *gorm.DB
}
